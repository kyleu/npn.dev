package routes

import (
	"github.com/kyleu/npn/npnweb"
	"net/http"
	"strings"

	"github.com/kyleu/npn/app/web/controllers"

	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/config"

	"github.com/gorilla/mux"
	"github.com/sagikazarmark/ocmux"
)

func BuildRouter(app *config.AppInfo) (*mux.Router, error) {
	npnweb.InitMime()

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

	// File System
	file := r.Path(p(util.KeyFile)).Subrouter()
	file.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.FileRoot))).Name(n(util.KeyFile, "root"))
	r.PathPrefix("/" + util.KeyFile + "/").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.FilePath))).Name(n(util.KeyFile))

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
	r.Path(p(util.KeySchema, "{key}", util.KeyModel, "{m}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SchemaModelDetail))).Name(n(util.KeySchema, util.KeyModel))

	// Project
	project := r.Path(p(util.KeyProject)).Subrouter()
	project.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ProjectList))).Name(n(util.KeyProject))
	r.Path(p(util.KeyProject, "new")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ProjectNew))).Name(n(util.KeyProject, "new"))
	r.Path(p(util.KeyProject, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ProjectDetail))).Name(n(util.KeyProject, "detail"))
	r.Path(p(util.KeyProject, "{key}", "edit")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ProjectEdit))).Name(n(util.KeyProject, "edit"))
	r.Path(p(util.KeyProject, "{key}", "edit")).Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.ProjectSave))).Name(n(util.KeyProject, "save"))
	r.Path(p(util.KeyProject, "{key}", "all")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TaskRunAll))).Name(n(util.KeyProject, util.KeyTask, "all"))
	r.Path(p(util.KeyProject, "{key}", "{task}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TaskRun))).Name(n(util.KeyProject, util.KeyTask))
	r.Path(p(util.KeyProject, "{key}", "{task}", "add")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TaskAdd))).Name(n(util.KeyProject, util.KeyTask, "add"))
	r.Path(p(util.KeyProject, "{key}", "{task}", "edit")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TaskEdit))).Name(n(util.KeyProject, util.KeyTask, "edit"))
	r.Path(p(util.KeyProject, "{key}", "{task}", "delete")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TaskDelete))).Name(n(util.KeyProject, util.KeyTask, "delete"))
	r.Path(p(util.KeyProject, "{key}", "task")).Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.TaskSave))).Name(n(util.KeyProject, util.KeyTask, "save"))

	// Sandbox
	sandbox := r.Path(p(util.KeySandbox)).Subrouter()
	sandbox.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SandboxList))).Name(n(util.KeySandbox))
	r.Path(p(util.KeySandbox, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SandboxRun))).Name(n(util.KeySandbox, "run"))

	// Routes
	routes := r.Path(p(util.KeyRoutes)).Subrouter()
	routes.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.RouteList))).Name(n(util.KeyRoutes))
	r.Path(p("sitemap.xml")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SitemapXML))).Name(n("sitemap"))
	r.Path(p(util.KeyModules)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ModuleList))).Name(n(util.KeyModules))

	// About
	about := r.Path(p(util.KeyAbout)).Subrouter()
	about.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.About))).Name(n(util.KeyAbout))

	// Assets
	_ = r.Path(p("assets")).Subrouter()
	r.Path(p("favicon.ico")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Favicon))).Name(n("favicon"))
	r.Path(p("robots.txt")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.RobotsTxt))).Name(n("robots"))
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
