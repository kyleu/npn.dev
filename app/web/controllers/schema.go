package controllers

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
	"net/http"
)

func SchemaList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx)
		schemata := ctx.App.Files.ListSchemata()
		return act.T(templates.SchemaList(schemata, ctx, w))
	})
}

func SchemaDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sch, err := ctx.App.Files.LoadSchema(key)
		if err != nil {
			return act.EResp(err, "cannot load schema")
		}
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx, "", sch.Key)
		return act.T(templates.SchemaDetail(sch, ctx, w))
	})
}

func SchemaEnumDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sch, err := ctx.App.Files.LoadSchema(key)
		if err != nil {
			return act.EResp(err, "cannot load schema")
		}
		e := mux.Vars(r)["e"]
		en := sch.Enums.Get(e)
		return act.T(templates.SchemaEnumDetail(sch, en, ctx, w))
	})
}

func SchemaModelDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sch, err := ctx.App.Files.LoadSchema(key)
		if err != nil {
			return act.EResp(err, "cannot load schema")
		}
		m := mux.Vars(r)["m"]
		model := sch.Models.Get(m)
		return act.T(templates.SchemaModelDetail(sch, model, ctx, w))
	})
}

func SchemaUnionDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sch, err := ctx.App.Files.LoadSchema(key)
		if err != nil {
			return act.EResp(err, "cannot load schema")
		}
		u := mux.Vars(r)["u"]
		union := sch.Unions.Get(u)
		return act.T(templates.SchemaUnionDetail(sch, union, ctx, w))
	})
}

func SchemaServiceDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sch, err := ctx.App.Files.LoadSchema(key)
		if err != nil {
			return act.EResp(err, "cannot load schema")
		}
		s := mux.Vars(r)["s"]
		svc := sch.Services.Get(s)
		return act.T(templates.SchemaServiceDetail(sch, svc, ctx, w))
	})
}

func schemaBreadcrumbs(ctx *web.RequestContext, pairs ...string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.KeySchema), util.KeySchema)
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, web.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}
