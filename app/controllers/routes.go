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
	collectionParam = "{c}"
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

	// Collections
	collection := r.Path(routes.Path("c")).Subrouter()
	collection.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(CollectionList))).Name(routes.Name("collection"))
	r.Path(routes.Path("c", "new")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(CollectionNew))).Name(routes.Name("collection", "new"))
	r.Path(routes.Path("c", collectionParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(CollectionDetail))).Name(routes.Name("collection", "detail"))
	r.Path(routes.Path("c", collectionParam, "edit")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(CollectionEdit))).Name(routes.Name("collection", "edit"))
	r.Path(routes.Path("c", collectionParam, "save")).Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(CollectionSave))).Name(routes.Name("collection", "save"))

	// Requests
	r.Path(routes.Path("c", "{c}", "new")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RequestNew))).Name(routes.Name("request", "new"))
	r.Path(routes.Path("c", "{c}", keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RequestDetail))).Name(routes.Name("request"))
	r.Path(routes.Path("c", "{c}", keyParam, "call")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RequestCall))).Name(routes.Name("request", "call"))
	r.Path(routes.Path("c", "{c}", keyParam, "edit")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RequestEdit))).Name(routes.Name("request", "edit"))
	r.Path(routes.Path("c", "{c}", keyParam, "transform")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RequestTransform))).Name(routes.Name("request", "transform"))
	r.Path(routes.Path("c", "{c}", keyParam, "delete")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RequestDelete))).Name(routes.Name("request", "delete"))

	// Ad hoc
	adhoc := r.Path(routes.Path("adhoc")).Subrouter()
	adhoc.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(AdhocForm))).Name(routes.Name("adhoc"))
	adhoc.Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(AdhocPost))).Name(routes.Name("adhoc", "post"))

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
