package npncontroller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npnuser"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"
)

// Entrypoint for most requests. It times the request, handles flashes, adds CORS headers, logs errors and handles redirects
func Act(w http.ResponseWriter, r *http.Request, f func(ctx *npnweb.RequestContext) (string, error)) {
	startNanos := npncore.TimerStart()
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

// An action that requires a successful auth.Check before proceeding
func AuthAct(w http.ResponseWriter, r *http.Request, f func(*npnweb.RequestContext) (string, error)) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		allowed := auth.Check(ctx.App.Auth(), ctx.Profile.UserID, ctx.Logger)
		if !allowed {
			const msg = "you are not authorized to see this page"
			if IsContentTypeJSON(GetContentType(r)) {
				ae := JSONResponse{Status: "error", Message: msg, Path: r.URL.Path, Occurred: time.Now()}
				return RespondJSON(w, "", ae, ctx.Logger)
			}
			return FlashAndRedir(false, msg, "home", w, r, ctx)
		}
		return f(ctx)
	})
}

// An action that requires a successful auth.Check and a user role of npnuser.RoleAdmin
func AdminAct(w http.ResponseWriter, r *http.Request, f func(*npnweb.RequestContext) (string, error)) {
	AuthAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if ctx.Profile.Role != npnuser.RoleAdmin {
			if IsContentTypeJSON(GetContentType(r)) {
				ae := JSONResponse{Status: "error", Message: "you are not an administrator", Path: r.URL.Path, Occurred: time.Now()}
				return RespondJSON(w, "", ae, ctx.Logger)
			}
			const msg = "you're not an administrator, silly!"
			return FlashAndRedir(false, msg, "home", w, r, ctx)
		}
		return f(ctx)
	})
}

// Creates breadcrumbs for admin actions
func AdminBC(ctx *npnweb.RequestContext, action string, name string) npnweb.Breadcrumbs {
	bc := npnweb.BreadcrumbsSimple(ctx.Route(npnweb.AdminLink()), npncore.KeyAdmin)
	bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(npnweb.AdminLink(action)), name)...)
	return bc
}
