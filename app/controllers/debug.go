package controllers

import (
	"net/http"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
)

func DebugList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Debug Stuff!"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("debug")}
		return npncontroller.T(templates.DebugList(ctx, w))
	})
}

func DebugRequest(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		u := r.URL.Query().Get("url")
		req := request.MockRequest
		if u != "" {
			p := request.PrototypeFromString(request.MethodGet, u)
			req = &request.Request{Key: "mock", Description: "ad-hoc request", Prototype: p}
		}
		result := app.Svc(ctx.App).Caller.Call(req.Prototype)
		ctx.Title = "Debug Request"
		ctx.Breadcrumbs = append(npnweb.BreadcrumbsSimple(ctx.Route("debug"), "debug"), npnweb.BreadcrumbSelf("request"))
		return npncontroller.T(templates.DebugCall(req, result, ctx, w))
	})
}
