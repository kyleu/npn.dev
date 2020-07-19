package controllers

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
)

func SchemaList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.PluralTitle(util.KeySchema)
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx)
		schemata, err := ctx.App.Schemata.Summaries()
		if err != nil {
			return act.EResp(err, "unable to load schemata")
		}
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
		err = ctx.App.Schemata.Save(nSch, true)
		if err != nil {
			return act.EResp(err, "unable to save schema from paths")
		}
		msg := fmt.Sprintf("Refreshed schema from [%v] paths", len(nSch.Paths))
		redir := ctx.Route(util.KeySchema+".detail", util.KeyKey, nSch.Key)
		return act.FlashAndRedir(true, msg, redir, w, r, ctx)
	})
}

func SchemaModelDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return act.EResp(err)
		}
		m := mux.Vars(r)["m"]
		pkg, key := util.SplitPackage(m)
		model := sch.Models.Get(pkg, key)
		if model == nil {
			return act.EResp(errors.New("cannot load model [" + m + "]"))
		}
		ctx.Title = model.Key
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx, ctx.Route(util.KeySchema+".detail", util.KeyKey, sch.Key), sch.Key, "", model.Key)
		return act.T(templates.SchemaModelDetail(sch, model, ctx, w))
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
	sch, err := ctx.App.Schemata.Load(key)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load schema ["+key+"]")
	}
	return sch, nil
}
