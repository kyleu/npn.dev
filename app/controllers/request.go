package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
)

const KeyRequest = "request"

func RequestDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll, req, err := loadRequest(r, ctx, "")
		if err != nil {
			return npncontroller.EResp(err)
		}
		return npncontroller.T(templates.RequestDetail(coll, req, nil, ctx, w))
	})
}

func RequestCall(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll, req, err := loadRequest(r, ctx, "call")
		if err != nil {
			return npncontroller.EResp(err)
		}
		result := app.Svc(ctx.App).Caller.Call(coll, req.Key, req.Prototype)
		return npncontroller.T(templates.CallDetail(coll, req, result, ctx, w))
	})
}

func RequestTransform(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_, req, err := loadRequest(r, ctx, "transform")
		if err != nil {
			return npncontroller.EResp(err)
		}
		return npncontroller.T(templates.RequestTransform(req.Prototype, ctx, w))
	})
}

func loadRequest(r *http.Request, ctx *npnweb.RequestContext, action string) (string, *request.Request, error) {
	c := mux.Vars(r)["c"]
	key := mux.Vars(r)[npncore.KeyKey]

	req, err := app.Svc(ctx.App).Collection.LoadRequest(c, key)
	if err != nil {
		return c, nil, err
	}

	ctx.Title = fmt.Sprintf("%v/%v", c, key)

	bc := append(
		npnweb.BreadcrumbsSimple(ctx.Route(KeyCollection), "collections"),
		npnweb.Breadcrumb{Path: ctx.Route(KeyCollection+".detail", "c", c), Title: c},
	)

	if len(action) == 0 {
		bc = append(bc, npnweb.BreadcrumbSelf(key))
	} else {
		ctx.Title = action + ": " + ctx.Title
		rt := ctx.Route(KeyRequest, "c", c, npncore.KeyKey, req.Key)
		bc = append(bc, npnweb.Breadcrumb{Path: rt, Title: req.Key}, npnweb.BreadcrumbSelf(action))
	}
	ctx.Breadcrumbs = bc

	return c, req, nil
}
