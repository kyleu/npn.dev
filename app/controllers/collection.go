package controllers

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
	"net/http"
)

const KeyCollection = "collection"

func CollectionList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		colls, err := app.Svc(ctx.App).Collection.List()
		if err != nil {
			return npncontroller.EResp(err)
		}
		ctx.Title = "Collections"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("collections")}
		return npncontroller.T(templates.CollectionList(colls, ctx, w))
	})
}

func CollectionDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)["c"]

		coll, err := app.Svc(ctx.App).Collection.Load(key)
		if err != nil {
			return npncontroller.EResp(err)
		}

		reqs, err := app.Svc(ctx.App).Collection.Requests(key)
		if err != nil {
			return npncontroller.EResp(err)
		}

		ctx.Title = "Collection"
		ctx.Breadcrumbs = append(npnweb.BreadcrumbsSimple(ctx.Route(KeyCollection), "collections"), npnweb.BreadcrumbSelf(key))
		return npncontroller.T(templates.CollectionDetail(coll, reqs, ctx, w))
	})
}
