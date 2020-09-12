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
	keyParam = "{key}"
)

func BuildRouter(app npnweb.AppInfo) (*mux.Router, error) {
	npncontroller.InitMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Home))).Name(routes.Name("home"))
	r.Path(routes.Path("health")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Health))).Name(routes.Name("health"))

	// Workspace
	workspace := r.Path(routes.Path("workspace")).Subrouter()
	workspace.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Workspace))).Name(routes.Name("workspace"))
	r.Path(routes.Path("s")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(Socket))).Name(routes.Name("websocket"))

	// Debug
	debug := r.Path(routes.Path("debug")).Subrouter()
	debug.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(DebugList))).Name(routes.Name("debug"))
	r.Path(routes.Path("debug", "request")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(DebugRequest))).Name(routes.Name("debug", "request"))

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
