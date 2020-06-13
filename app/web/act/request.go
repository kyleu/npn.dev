package act

import (
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/kyleu/npn/app/web"
)

func GetContentType(r *http.Request) string {
	ret := r.Header.Get("content-type")
	idx := strings.Index(ret, ";")
	if idx > -1 {
		ret = ret[0:idx]
	}
	return strings.TrimSpace(ret)
}

func IsContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}

func SaveSession(w http.ResponseWriter, r *http.Request, ctx *web.RequestContext) {
	ctx.Session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteDefaultMode}
	err := ctx.Session.Save(r, w)
	if err != nil {
		ctx.Logger.Warn("unable to save session to response")
	}
}

func FlashAndRedir(success bool, msg string, redir string, w http.ResponseWriter, r *http.Request, ctx *web.RequestContext) (string, error) {
	status := "error"
	if success {
		status = "success"
	}
	ctx.Session.AddFlash(status + ":" + msg)
	SaveSession(w, r, ctx)
	if strings.HasPrefix(redir, "/") {
		return redir, nil
	}
	if strings.HasPrefix(redir, "http") {
		ctx.Logger.Warn("flash redirect attempted for non-local request")
		return "/", nil
	}
	return ctx.Route(redir), nil
}
