package controllers

import (
	"fmt"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/gen/templates"
)

func SchemaList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.PluralTitle("schema")
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx)
		schemata, err := app.Schemata(ctx.App).Summaries()
		if err != nil {
			return npncontroller.EResp(err, "unable to load schemata")
		}
		return npncontroller.T(templates.SchemaList(schemata, ctx, w))
	})
}

func SchemaDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return npncontroller.EResp(err)
		}
		ctx.Title = sch.Title
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx, "", sch.Key)
		return npncontroller.T(templates.SchemaDetail(sch, ctx, w))
	})
}

func SchemaRefresh(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return npncontroller.EResp(err)
		}
		nSch, err := app.Parsers(ctx.App).Refresh(sch)
		if err != nil {
			return npncontroller.EResp(err, "error loading schema from paths")
		}
		err = app.Schemata(ctx.App).Save(nSch, true)
		if err != nil {
			return npncontroller.EResp(err, "unable to save schema from paths")
		}
		msg := fmt.Sprintf("Refreshed schema from [%v] paths", len(nSch.Paths))
		redir := ctx.Route("schema"+".detail", npncore.KeyKey, nSch.Key)
		return npncontroller.FlashAndRedir(true, msg, redir, w, r, ctx)
	})
}

func SchemaModelDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		sch, err := schemaFromRequest(ctx, r)
		if err != nil {
			return npncontroller.EResp(err)
		}
		m := mux.Vars(r)["m"]
		pkg, key := util.SplitPackage(m)
		model := sch.Models.Get(pkg, key)
		if model == nil {
			return npncontroller.EResp(errors.New("cannot load model [" + m + "]"))
		}
		ctx.Title = model.Key
		ctx.Breadcrumbs = schemaBreadcrumbs(ctx, ctx.Route("schema"+".detail", npncore.KeyKey, sch.Key), sch.Key, "", model.Key)
		return npncontroller.T(templates.SchemaModelDetail(sch, model, ctx, w))
	})
}

func schemaBreadcrumbs(ctx *npnweb.RequestContext, pairs ...string) npnweb.Breadcrumbs {
	bc := npnweb.BreadcrumbsSimple(ctx.Route("schema"), "schema")
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, npnweb.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}

func schemaFromRequest(ctx *npnweb.RequestContext, r *http.Request) (*schema.Schema, error) {
	key := mux.Vars(r)[npncore.KeyKey]
	sch, err := app.Schemata(ctx.App).Load(key)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load schema ["+key+"]")
	}
	return sch, nil
}
