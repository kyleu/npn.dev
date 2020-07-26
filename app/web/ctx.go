package web

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kyleu/npn/app/config"
	"logur.dev/logur"
)

type RequestContext struct {
	App         *config.AppInfo
	Logger      logur.Logger
	Profile     *npnuser.UserProfile
	Routes      *mux.Router
	Request     *url.URL
	Title       string
	Breadcrumbs npnweb.Breadcrumbs
	Flashes     []string
	Session     *sessions.Session
}

func (r *RequestContext) Route(act string, pairs ...string) string {
	return npnweb.Route(r.Routes, r.Logger, act, pairs...)
}

func ExtractContext(w http.ResponseWriter, r *http.Request) *RequestContext {
	ai, ok := r.Context().Value(npncore.InfoKey).(*config.AppInfo)
	if !ok {
		ai.Logger.Warn("cannot load AppInfo")
	}
	routes, ok := r.Context().Value(npncore.RoutesKey).(*mux.Router)
	if !ok {
		ai.Logger.Warn("cannot load Router")
	}
	session, err := npnweb.Store.Get(r, npncore.AppName + "-session")
	if err != nil {
		session = sessions.NewSession(npnweb.Store, npncore.AppName + "-session")
	}

	logger := logur.WithFields(ai.Logger, map[string]interface{}{"path": r.URL.Path, "method": r.Method})

	prof, err := loadProfile(ai.Files)
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
		Title:       npncore.AppName,
		Breadcrumbs: nil,
		Flashes:     flashes,
		Session:     session,
	}
}

func loadProfile(f *npncore.FileLoader) (*npnuser.UserProfile, error) {
	content, err := f.ReadFile("profile.json")
	if err != nil {
		return npnuser.NewUserProfile(), nil
	}
	tgt := &npnuser.UserProfile{}
	npncore.FromJSON([]byte(content), tgt, nil)
	return tgt, nil
}

