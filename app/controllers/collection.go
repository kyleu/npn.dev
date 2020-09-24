package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
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

func CollectionNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll := &collection.Collection{}
		ctx.Title = "New collection"
		ctx.Breadcrumbs = append(npnweb.BreadcrumbsSimple(ctx.Route(KeyCollection), "collections"), npnweb.BreadcrumbSelf("new"))
		return npncontroller.T(templates.CollectionForm("new", coll, true, ctx, w))
	})
}

func CollectionDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)["c"]

		coll, err := app.Svc(ctx.App).Collection.Load(key)
		if err != nil {
			return npncontroller.EResp(err)
		}

		reqs, err := app.Svc(ctx.App).Collection.ListRequests(key)
		if err != nil {
			return npncontroller.EResp(err)
		}

		ctx.Title = "Collection"
		ctx.Breadcrumbs = append(npnweb.BreadcrumbsSimple(ctx.Route(KeyCollection), "collections"), npnweb.BreadcrumbSelf(key))
		return npncontroller.T(templates.CollectionDetail(coll, reqs, ctx, w))
	})
}

func CollectionEdit(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)["c"]

		coll, err := app.Svc(ctx.App).Collection.Load(key)
		if err != nil {
			return npncontroller.EResp(err)
		}

		ctx.Title = coll.Title
		bc := npnweb.Breadcrumb{Path: ctx.Route(KeyCollection+".detail", "c", key), Title: key}
		ctx.Breadcrumbs = append(npnweb.BreadcrumbsSimple(ctx.Route(KeyCollection), "collections"), bc, npnweb.BreadcrumbSelf("edit"))
		return npncontroller.T(templates.CollectionForm(coll.Key, coll, false, ctx, w))
	})
}

func CollectionSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		originalKey := mux.Vars(r)["c"]

		_ = r.ParseForm()
		key := r.Form.Get("key")
		if len(key) == 0 {
			key = originalKey
		}
		title := r.Form.Get("title")
		description := r.Form.Get("description")

		err := app.Svc(ctx.App).Collection.Save(originalKey, key, title, description)
		if err != nil {
			return npncontroller.EResp(err)
		}

		return ctx.Route(KeyCollection+".detail", "c", key), nil
	})
}

func CollectionDelete(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll := mux.Vars(r)["c"]

		err := app.Svc(ctx.App).Collection.Delete(coll)
		if err != nil {
			return npncontroller.EResp(err, "unable to delete collection ["+coll+"]")
		}

		msg := "deleted collection [" + coll + "]"
		return npncontroller.FlashAndRedir(true, msg, ctx.Route(KeyCollection), w, r, ctx)
	})
}
