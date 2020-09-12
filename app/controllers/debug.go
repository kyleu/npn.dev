package controllers

import (
	"github.com/kyleu/npn/npncore"
	"net/http"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
)

func DebugList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		files := ctx.App.Files().ListJSON("collections/debug/requests")

		ctx.Title = "Debug Stuff!"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("debug")}
		return npncontroller.T(templates.DebugList(files, ctx, w))
	})
}

func DebugRequest(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		svc := app.Svc(ctx.App)

		req := request.MockRequest

		k := r.URL.Query().Get("key")
		if k != "" {
			fn := "collections/debug/requests/" + k + ".json"
			content, err := ctx.App.Files().ReadFile(fn)
			if err != nil {
				return npncontroller.EResp(err)
			}
			r := request.Request{}
			err = npncore.FromJSON([]byte(content), &r)
			if err != nil {
				return npncontroller.EResp(err)
			}
			req = r.Normalize(k)
		}

		u := r.URL.Query().Get("url")
		if u != "" {
			p := request.PrototypeFromString(request.MethodGet, u)
			req = &request.Request{Key: "mock", Description: "ad hoc request", Prototype: p}
		}

		result := svc.Caller.Call(req.Prototype)

		ctx.Title = "Debug Request"
		ctx.Breadcrumbs = append(npnweb.BreadcrumbsSimple(ctx.Route("debug"), "debug"), npnweb.BreadcrumbSelf("request"))
		return npncontroller.T(templates.DebugCall(req, result, ctx, w))
	})
}
