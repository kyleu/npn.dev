package act

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"net/http"
	"strings"
	"time"

	"github.com/kyleu/npn/gen/components"

	"emperror.dev/errors"
	"logur.dev/logur"

	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"golang.org/x/text/language"
)

type errorResult struct {
	Status  string
	Message string
}

func Act(w http.ResponseWriter, r *http.Request, f func(*web.RequestContext) (string, error)) {
	startNanos := time.Now().UnixNano()
	ctx := web.ExtractContext(w, r)

	if len(ctx.Flashes) > 0 {
		SaveSession(w, r, ctx)
	}

	redir, err := f(ctx)
	if err != nil {
		ctx.Logger.Warn(fmt.Sprintf("error running action: %+v", err))
		if len(ctx.Title) == 0 {
			ctx.Title = "Error"
		}
		if IsContentTypeJSON(GetContentType(r)) {
			_, _ = RespondJSON(w, "", errorResult{Status: util.KeyError, Message: err.Error()}, ctx.Logger)
		} else {
			_, _ = components.InternalServerError(npncore.GetErrorDetail(err), r, ctx, w)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if len(ba) == 0 {
		return "", errors.New("no bytes available to write")
	}
	_, err := w.Write(ba)
	return "", errors.Wrap(err, "cannot write to response")
}

func logComplete(startNanos int64, ctx *web.RequestContext, status int, r *http.Request) {
	delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	ms := npncore.MicrosToMillis(language.AmericanEnglish, int(delta))
	args := map[string]interface{}{"elapsed": delta, util.KeyStatus: status}
	msg := fmt.Sprintf("[%v %v] returned [%v] in [%v]", r.Method, r.URL.Path, status, ms)
	ctx.Logger.Debug(msg, args)
}
