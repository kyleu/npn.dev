package npncontroller

import (
	"net/http"
	"strings"

	"github.com/kyleu/npn/npncore"

	"github.com/kyleu/npn/npnweb"

	"github.com/gorilla/sessions"
)

// Returns the content type of the request
func GetContentType(r *http.Request) string {
	ret := r.Header.Get("Content-Type")
	idx := strings.Index(ret, ";")
	if idx > -1 {
		ret = ret[0:idx]
	}
	return strings.TrimSpace(ret)
}

// Returns true if the type is in fact JSON
func IsContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}

// Saves the session, call after mutating
func SaveSession(w http.ResponseWriter, r *http.Request, ctx *npnweb.RequestContext) {
	ctx.Session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteDefaultMode}
	err := ctx.Session.Save(r, w)
	if err != nil {
		ctx.Logger.Warn("unable to save session to response")
	}
}

// Sets a flash cookie (which is handled elsewhere) and redirects to a url
func FlashAndRedir(success bool, msg string, redir string, w http.ResponseWriter, r *http.Request, ctx *npnweb.RequestContext) (string, error) {
	status := npncore.KeyError
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
