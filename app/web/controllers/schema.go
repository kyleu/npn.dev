package controllers

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
	"net/http"
)

func SchemaList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.PluralTitle(util.KeySchema)
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx)
		schemata := ctx.App.Files.ListSchemata()
		return act.T(templates.SchemaList(schemata, ctx, w))
	})
}

func SchemaDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return act.EResp(err)
		}
		ctx.Title = sch.Title
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx, "", sch.Key)
		return act.T(templates.SchemaDetail(sch, ctx, w))
	})
}

func SchemaRefresh(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return act.EResp(err)
		}
		nSch, err := ctx.App.Parsers.Refresh(sch)
		if err != nil {
			return act.EResp(err, "error loading schema from paths")
		}
		err = ctx.App.Files.SaveSchema(nSch, true)
		if err != nil {
			return act.EResp(err, "unable to save schema from paths")
		}
		msg := fmt.Sprintf("Refreshed schema from [%v] paths", len(nSch.Paths))
		redir := ctx.Route(util.KeySchema + ".detail", util.KeyKey, nSch.Key)
		return act.FlashAndRedir(true, msg, redir, w, r, ctx)
	})
}

func SchemaEnumDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return act.EResp(err)
		}
		e := mux.Vars(r)["e"]
		en := sch.Enums.Get(e)
		if en == nil {
			return act.EResp(err, "cannot load enum [" + e + "]")
		}
		ctx.Title = en.Key
		return act.T(templates.SchemaEnumDetail(sch, en, ctx, w))
	})
}

func SchemaModelDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return act.EResp(err)
		}
		m := mux.Vars(r)["m"]
		model := sch.Models.Get(m)
		if model == nil {
			return act.EResp(err, "cannot load model [" + m + "]")
		}
		ctx.Title = util.PluralTitle(util.KeySchema)
		return act.T(templates.SchemaModelDetail(sch, model, ctx, w))
	})
}

func SchemaUnionDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return act.EResp(err)
		}
		u := mux.Vars(r)["u"]
		union := sch.Unions.Get(u)
		if union == nil {
			return act.EResp(err, "cannot load union [" + u + "]")
		}
		ctx.Title = util.PluralTitle(util.KeySchema)
		return act.T(templates.SchemaUnionDetail(sch, union, ctx, w))
	})
}

func schemaBreadcrumbs(ctx *web.RequestContext, pairs ...string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.KeySchema), util.KeySchema)
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, web.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}

func schemaFromRequest(ctx *web.RequestContext, r *http.Request) (*schema.Schema, error) {
	key := mux.Vars(r)[util.KeyKey]
	sch, err := ctx.App.Files.LoadSchema(key)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load schema [" + key + "]")
	}
	return sch, nil
}
