package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npnasset"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/sagikazarmark/ocmux"
)

const (
	keyParam  = "{key}"
	taskParam = "{task}"
)

func BuildRouter(app npnweb.AppInfo) (*mux.Router, error) {
	npncontroller.InitMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Home))).Name(routes.Name("home"))
	r.Path(routes.Path("health")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Health))).Name(routes.Name("health"))

	// DataSource
	dsn := r.Path(routes.Path("datasource")).Subrouter()
	dsn.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(DataSourceList))).Name(routes.Name("datasource"))
	r.Path(routes.Path("datasource", "{t}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(DataSourceDetail))).Name(routes.Name("datasource", "detail"))
	r.Path(routes.Path("datasource", "{t}")).Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(DataSourceSave))).Name(routes.Name("datasource", "save"))

	// Schema
	schema := r.Path(routes.Path("schema")).Subrouter()
	schema.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SchemaList))).Name(routes.Name("schema"))
	r.Path(routes.Path("schema", keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SchemaDetail))).Name(routes.Name("schema", "detail"))
	r.Path(routes.Path("schema", keyParam, "refresh")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SchemaRefresh))).Name(routes.Name("schema", "refresh"))
	r.Path(routes.Path("schema", keyParam, "model", "{m}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SchemaModelDetail))).Name(routes.Name("schema", "model"))

	// Project
	project := r.Path(routes.Path("project")).Subrouter()
	project.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectList))).Name(routes.Name("project"))
	r.Path(routes.Path("project", "new")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectNew))).Name(routes.Name("project", "new"))
	r.Path(routes.Path("project", keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectDetail))).Name(routes.Name("project", "detail"))
	r.Path(routes.Path("project", keyParam, "edit")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectEdit))).Name(routes.Name("project", "edit"))
	r.Path(routes.Path("project", keyParam, "edit")).Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectSave))).Name(routes.Name("project", "save"))
	r.Path(routes.Path("project", keyParam, "models")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectModels))).Name(routes.Name("project", "models"))
	r.Path(routes.Path("project", keyParam, "models")).Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(ProjectModelsSave))).Name(routes.Name("project", "models", "save"))
	r.Path(routes.Path("project", keyParam, "all")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(TaskRunAll))).Name(routes.Name("project", "task", "all"))
	r.Path(routes.Path("project", keyParam, taskParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(TaskRun))).Name(routes.Name("project", "task"))
	r.Path(routes.Path("project", keyParam, taskParam, "add")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(TaskAdd))).Name(routes.Name("project", "task", "add"))
	r.Path(routes.Path("project", keyParam, taskParam, "edit")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(TaskEdit))).Name(routes.Name("project", "task", "edit"))
	r.Path(routes.Path("project", keyParam, taskParam, "delete")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(TaskDelete))).Name(routes.Name("project", "task", "delete"))
	r.Path(routes.Path("project", keyParam, "task")).Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(TaskSave))).Name(routes.Name("project", "task", "save"))

	// Sandbox
	sandbox := r.Path(routes.Path(npncore.KeySandbox)).Subrouter()
	sandbox.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SandboxList))).Name(routes.Name(npncore.KeySandbox))
	r.Path(routes.Path(npncore.KeySandbox, keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SandboxRun))).Name(routes.Name(npncore.KeySandbox, "run"))

	// About
	about := r.Path(routes.Path(npncore.KeyAbout)).Subrouter()
	about.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(About))).Name(routes.Name(npncore.KeyAbout))

	// Assets
	_ = r.Path(routes.Path("assets")).Subrouter()
	r.Path(routes.Path("favicon.ico")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Favicon))).Name(routes.Name("favicon"))
	r.Path(routes.Path("robots.txt")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RobotsTxt))).Name(routes.Name("robots"))
	r.PathPrefix(routes.Path("vendor")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(npnasset.VendorAsset))).Name(routes.Name("vendor"))
	r.PathPrefix(routes.Path("assets")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Static))).Name(routes.Name("assets"))

	// Provided
	npncontroller.RoutesProfile(app, r)
	npncontroller.RoutesFile(app, r, func(path string) string {
		return ""
	})
	npncontroller.RoutesUtil(app, r)

	return r, nil
}
