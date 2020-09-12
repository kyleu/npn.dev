package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"
	"net/http"
)

const KeyRequest = "request"

func AdhocForm(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Ad hoc Request"
		ctx.Breadcrumbs = npnweb.BreadcrumbsSimple("", "ad hoc")
		return npncontroller.T(npntemplate.StaticMessage("TODO", ctx, w))
	})
}

func AdhocPost(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Ad hoc Post"
		ctx.Breadcrumbs = npnweb.BreadcrumbsSimple("", "ad hoc")
		return npncontroller.T(npntemplate.StaticMessage("TODO", ctx, w))
	})
}

func RequestNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll := mux.Vars(r)["c"]
		req := request.NewRequest()
		return npncontroller.T(templates.RequestForm(coll, req, nil, ctx, w))
	})
}

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
		result := app.Svc(ctx.App).Caller.Call(req.Prototype)
		return npncontroller.T(templates.CallDetail(coll, req, result, ctx, w))
	})
}

func RequestEdit(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll, req, err := loadRequest(r, ctx, "edit")
		if err != nil {
			return npncontroller.EResp(err)
		}
		return npncontroller.T(templates.RequestForm(coll, req, nil, ctx, w))
	})
}

func RequestDelete(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll := mux.Vars(r)["c"]
		key := mux.Vars(r)[npncore.KeyKey]
		msg := "deleted request [" + key + "] from this collection"
		return npncontroller.FlashAndRedir(true, msg, ctx.Route(KeyCollection+".detail", "c", coll), w, r, ctx)
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
