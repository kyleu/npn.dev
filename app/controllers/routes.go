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
	keyParam        = "{key}"
	collectionParam = "{c}"
)

func BuildRouter(ai npnweb.AppInfo) (*mux.Router, error) {
	npncontroller.InitMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Path(routes.Path("system")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(System))).Name(routes.Name("system"))
	r.Path(routes.Path("health")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Health))).Name(routes.Name("health"))

	// Workspace
	r.Path("/").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("home"))
	r.PathPrefix(routes.Path("c/")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("home", "collection"))
	r.PathPrefix(routes.Path("r/")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Workspace))).Name(routes.Name("home", "result"))

	r.Path(routes.Path("svg", "gantt")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(npncontroller.Gantt))).Name(routes.Name("svg", "gantt"))
	r.Path(routes.Path("s")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Socket))).Name(routes.Name("websocket"))

	// Import
	imprt := r.Path(routes.Path("i")).Subrouter()
	imprt.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(ImportForm))).Name(routes.Name("import", "form"))
	imprt.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(ImportUpload))).Name(routes.Name("import", "upload"))
	r.Path(routes.Path("i", keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(ImportDetail))).Name(routes.Name("import", "detail"))

	// Collections
	collection := r.Path(routes.Path("browse")).Subrouter()
	collection.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(CollectionList))).Name(routes.Name(npncore.KeyCollection))
	r.Path(routes.Path("browse", "new")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(CollectionNew))).Name(routes.Name(npncore.KeyCollection, "new"))
	collectionPath := routes.Path("browse", collectionParam)
	r.Path(collectionPath).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(CollectionDetail))).Name(routes.Name(npncore.KeyCollection, "detail"))
	r.Path(collectionPath + "/act/edit").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(CollectionEdit))).Name(routes.Name(npncore.KeyCollection, "edit"))
	r.Path(collectionPath + "/act/save").Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(CollectionSave))).Name(routes.Name(npncore.KeyCollection, "save"))
	r.Path(collectionPath + "/act/delete").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(CollectionDelete))).Name(routes.Name(npncore.KeyCollection, "delete"))

	// Requests
	r.Path(routes.Path("browse", collectionParam, "new")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestNew))).Name(routes.Name("request", "new"))
	requestPath := collectionPath + "/" + keyParam
	r.Path(requestPath).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestDetail))).Name(routes.Name("request"))
	r.Path(requestPath + "/call").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestCall))).Name(routes.Name("request", "call"))
	r.Path(requestPath + "/edit").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestEdit))).Name(routes.Name("request", "edit"))
	r.Path(requestPath + "/save").Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestSave))).Name(routes.Name("request", "save"))
	r.Path(requestPath + "/transform").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestTransform))).Name(routes.Name("request", "transform"))
	r.Path(requestPath + "/delete").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RequestDelete))).Name(routes.Name("request", "delete"))

	// Test
	test := r.Path(routes.Path(npncore.KeyTest)).Subrouter()
	test.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(TestIndex))).Name(routes.Name(npncore.KeyTest))
	r.PathPrefix(routes.Path("test/")).Handler(routes.AddContext(r, ai, http.StripPrefix(routes.Path(npncore.KeyTest), http.HandlerFunc(TestCall)))).Name(routes.Name(npncore.KeyTest, "run"))

	// Sandbox
	sandbox := r.Path(routes.Path(npncore.KeySandbox)).Subrouter()
	sandbox.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(SandboxList))).Name(routes.Name(npncore.KeySandbox))
	r.Path(routes.Path(npncore.KeySandbox, keyParam)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(SandboxRun))).Name(routes.Name(npncore.KeySandbox, "run"))

	// About
	about := r.Path(routes.Path(npncore.KeyAbout)).Subrouter()
	about.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(About))).Name(routes.Name(npncore.KeyAbout))

	// Assets
	_ = r.Path(routes.Path("assets")).Subrouter()
	r.Path(routes.Path("favicon.ico")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Favicon))).Name(routes.Name("favicon"))
	r.Path(routes.Path("robots.txt")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(RobotsTxt))).Name(routes.Name("robots"))
	r.PathPrefix(routes.Path("vendor")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(npnasset.VendorAsset))).Name(routes.Name("vendor"))
	r.PathPrefix(routes.Path("assets")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(Static))).Name(routes.Name("assets"))

	r.PathPrefix(routes.Adm("src/")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.StripPrefix(routes.Adm("src/"), http.HandlerFunc(Source)))).Name(npnweb.AdminLink("source"))

	// Provided
	npncontroller.RoutesProfile(ai, r)
	npncontroller.RoutesFile(ai, r)
	npncontroller.RoutesUtil(ai, r)

	return r, nil
}
