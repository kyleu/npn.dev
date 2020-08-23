package npncontroller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kyleu/npn/npnuser"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"

	"emperror.dev/errors"
	"logur.dev/logur"

	"golang.org/x/text/language"
)

type errorResult struct {
	Status  string
	Message string
}

func Act(w http.ResponseWriter, r *http.Request, f func(ctx *npnweb.RequestContext) (string, error)) {
	startNanos := time.Now().UnixNano()
	ctx := npnweb.ExtractContext(w, r, false)

	if len(ctx.Flashes) > 0 {
		SaveSession(w, r, ctx)
	}

	WriteCORS(w)

	redir, err := f(ctx)
	if err != nil {
		ctx.Logger.Warn(fmt.Sprintf("error running action: %+v", err))
		if len(ctx.Title) == 0 {
			ctx.Title = "Error"
		}
		if IsContentTypeJSON(GetContentType(r)) {
			_, _ = RespondJSON(w, "", errorResult{Status: npncore.KeyError, Message: err.Error()}, ctx.Logger)
		} else {
			_, _ = npntemplate.InternalServerError(npncore.GetErrorDetail(err), r, ctx, w)
		}
	}
	if redir != "" {
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
		logComplete(startNanos, ctx, http.StatusFound, r)
	} else {
		logComplete(startNanos, ctx, http.StatusOK, r)
	}
}

func T(_ int, err error) (string, error) {
	return "", err
}

func EResp(err error, msgs ...string) (string, error) {
	msg := strings.Join(msgs, "\n")
	if len(msg) == 0 {
		return "", err
	}
	return "", errors.Wrap(err, msg)
}

func ENew(msg string) (string, error) {
	return "", errors.New(msg)
}

func RespondJSON(w http.ResponseWriter, filename string, body interface{}, logger logur.Logger) (string, error) {
	return RespondMIME(filename, "application/json", "pdf", npncore.ToJSONBytes(body, logger), w)
}

func RespondMIME(filename string, mime string, ext string, ba []byte, w http.ResponseWriter) (string, error) {
	w.Header().Set("Content-Type", mime+"; charset=UTF-8")
	if len(filename) > 0 {
		if !strings.HasSuffix(filename, "."+ext) {
			filename = filename + "." + ext
		}
		w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	}
	WriteCORS(w)
	if len(ba) == 0 {
		return "", errors.New("no bytes available to write")
	}
	_, err := w.Write(ba)
	return "", errors.Wrap(err, "cannot write to response")
}

func WriteCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

type JSONResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Path     string    `json:"path"`
	Occurred time.Time `json:"occurred"`
}

func AdminAct(w http.ResponseWriter, r *http.Request, f func(*npnweb.RequestContext) (string, error)) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if ctx.Profile.Role != npnuser.RoleAdmin {
			if IsContentTypeJSON(GetContentType(r)) {
				ae := JSONResponse{Status: "error", Message: "you are not an administrator", Path: r.URL.Path, Occurred: time.Now()}
				return RespondJSON(w, "", ae, ctx.Logger)
			}
			msg := "you're not an administrator, silly!"
			return FlashAndRedir(false, msg, "home", w, r, ctx)
		}
		return f(ctx)
	})
}

func AdminBC(ctx *npnweb.RequestContext, action string, name string) npnweb.Breadcrumbs {
	bc := npnweb.BreadcrumbsSimple(ctx.Route(npnweb.AdminLink()), npncore.KeyAdmin)
	bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(npnweb.AdminLink(action)), name)...)
	return bc
}

func logComplete(startNanos int64, ctx *npnweb.RequestContext, status int, r *http.Request) {
	delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	ms := npncore.MicrosToMillis(language.AmericanEnglish, int(delta))
	args := map[string]interface{}{"elapsed": delta, npncore.KeyStatus: status}
	msg := fmt.Sprintf("[%v %v] returned [%v] in [%v]", r.Method, r.URL.Path, status, ms)
	ctx.Logger.Debug(msg, args)
}
