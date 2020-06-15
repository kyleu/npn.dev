package routes

import (
	"net/http"
	"strings"

	"github.com/kyleu/npn/app/web/controllers"

	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/config"

	"github.com/gorilla/mux"
	"github.com/sagikazarmark/ocmux"
)

func BuildRouter(app *config.AppInfo) (*mux.Router, error) {
	initMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Home))).Name(n("home"))
	r.Path(p("health")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Health))).Name(n("health"))

	// Profile
	profile := r.Path(p(util.KeyProfile)).Subrouter()
	profile.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Profile))).Name(n(util.KeyProfile))
	profile.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.ProfileSave))).Name(n(util.KeyProfile, "save"))
	r.Path(p(util.KeyProfile, "theme", "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ProfileTheme))).Name(n(util.KeyProfile, util.KeyTheme))

	// DataSource
	dsn := r.Path(p(util.KeyDataSource)).Subrouter()
	dsn.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.DataSourceList))).Name(n(util.KeyDataSource))
	r.Path(p(util.KeyDataSource, "{t}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.DataSourceDetail))).Name(n(util.KeyDataSource, "detail"))
	r.Path(p(util.KeyDataSource, "{t}")).Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.DataSourceSave))).Name(n(util.KeyDataSource, "save"))

	// Schema
	schema := r.Path(p(util.KeySchema)).Subrouter()
	schema.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaList))).Name(n(util.KeySchema))
	r.Path(p(util.KeySchema, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaDetail))).Name(n(util.KeySchema, "detail"))
	r.Path(p(util.KeySchema, "{key}", "refresh")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaRefresh))).Name(n(util.KeySchema, "refresh"))
	r.Path(p(util.KeySchema, "{key}", util.KeyEnum, "{e}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaEnumDetail))).Name(n(util.KeySchema, util.KeyEnum))
	r.Path(p(util.KeySchema, "{key}", util.KeyModel, "{m}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaModelDetail))).Name(n(util.KeySchema, util.KeyModel))
	r.Path(p(util.KeySchema, "{key}", util.KeyUnion, "{u}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaUnionDetail))).Name(n(util.KeySchema, util.KeyUnion))

	// Sandbox
	r.Path(p(util.KeySandbox)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SandboxList))).Name(n(util.KeySandbox))
	r.Path(p(util.KeySandbox, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SandboxRun))).Name(n(util.KeySandbox, "run"))

	// About
	r.Path(p(util.KeyAbout)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.About))).Name(n(util.KeyAbout))

	// Assets
	r.Path(p("favicon.ico")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Favicon))).Name(n("favicon"))
	r.Path(p("robots.txt")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.RobotsTxt))).Name(n("robots"))
	r.Path(p("sitemap.xml")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SitemapXML))).Name(n("sitemap"))
	r.PathPrefix(p("assets")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Static))).Name(n("assets"))

	r.PathPrefix("").Handler(addContext(r, app, http.HandlerFunc(controllers.NotFound)))

	return r, nil
}

func p(params ...string) string {
	ret := ""
	for _, p := range params {
		ret = ret + "/" + p
	}
	return ret
}

func n(params ...string) string {
	return strings.Join(params, ".")
}
