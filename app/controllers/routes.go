package controllers

import (
	"net/http"

	"github.com/kyleu/libnpn/npncore"

	"github.com/gorilla/mux"
	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncontroller/routes"
	"github.com/kyleu/libnpn/npnweb"
	"github.com/sagikazarmark/ocmux"
)

const (
	keyParam = "{key}"
)

func BuildRouter(ai npnweb.AppInfo) (*mux.Router, error) {
	npncontroller.InitMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(WorkspaceIndex))).Name(routes.Name("home"))
	r.Path(routes.Path("health")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Health))).Name(routes.Name("health"))

	// Static
	r.Path(routes.Path("download")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Download))).Name(routes.Name("download"))

	// Export
	r.PathPrefix(routes.Path("x")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("export"))

	// Utility
	r.Path(routes.Path("about")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("about"))
	r.Path(routes.Path("cfg")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("config"))
	r.Path(routes.Path("help")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("help"))
	r.Path(routes.Path("u")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("settings"))

	// Session
	r.Path(routes.Path("s")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("sessions"))
	r.PathPrefix(routes.Path("s/")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("session"))

	// Collection
	r.Path(routes.Path("c")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("collections"))
	r.PathPrefix(routes.Path("c/")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("collection"))

	// Response
	r.PathPrefix(routes.Path("r")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("result"))
	r.Path(routes.Path("svg", "gantt")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Gantt))).Name(routes.Name("svg", "gantt"))

	// Socket
	r.Path(routes.Path("ws")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Socket))).Name(routes.Name("websocket"))

	// Import
	imprt := r.Path(routes.Path("i")).Subrouter()
	imprt.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name(npncore.KeyImport, "form"))
	imprt.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(ImportUpload))).Name(routes.Name(npncore.KeyImport, "upload"))
	r.Path(routes.Path("i", keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name(npncore.KeyImport, "detail"))

	// Assets
	_ = r.Path(routes.Path("assets")).Subrouter()
	r.Path(routes.Path("favicon.ico")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Favicon))).Name(routes.Name("favicon"))
	r.Path(routes.Path("robots.txt")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RobotsTxt))).Name(routes.Name("robots"))
	r.PathPrefix(routes.Path("assets")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Static))).Name(routes.Name("assets"))

	// Provided
	npncontroller.RoutesAuth(ai, r)
	npncontroller.RoutesUtil(ai, r)

	return r, nil
}
