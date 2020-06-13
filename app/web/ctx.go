package web

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kyleu/npn/app/config"
	"github.com/kyleu/npn/app/util"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"logur.dev/logur"
)

type RequestContext struct {
	App         *config.AppInfo
	Logger      logur.Logger
	Profile     *util.UserProfile
	Routes      *mux.Router
	Request     *url.URL
	Title       string
	Breadcrumbs Breadcrumbs
	Flashes     []string
	Session     *sessions.Session
}

func (r *RequestContext) Route(act string, pairs ...string) string {
	return Route(r.Routes, r.Logger, act, pairs...)
}

func ExtractContext(w http.ResponseWriter, r *http.Request) *RequestContext {
	ai, ok := r.Context().Value(util.InfoKey).(*config.AppInfo)
	if !ok {
		ai.Logger.Warn("cannot load AppInfo")
	}
	routes, ok := r.Context().Value(util.RoutesKey).(*mux.Router)
	if !ok {
		ai.Logger.Warn("cannot load Router")
	}
	session, err := store.Get(r, sessionName)
	if err != nil {
		session = sessions.NewSession(store, sessionName)
	}

	logger := logur.WithFields(ai.Logger, map[string]interface{}{"path": r.URL.Path, "method": r.Method})

	prof, err := ai.Files.LoadProfile()
	if err != nil {
		logger.Warn(fmt.Sprintf("unable to load profile: %+v", err))
	}

	flashes := make([]string, 0)
	for _, f := range session.Flashes() {
		flashes = append(flashes, fmt.Sprint(f))
	}

	return &RequestContext{
		App:         ai,
		Logger:      logger,
		Profile:     prof,
		Routes:      routes,
		Request:     r.URL,
		Title:       util.AppName,
		Breadcrumbs: nil,
		Flashes:     flashes,
		Session:     session,
	}
}
