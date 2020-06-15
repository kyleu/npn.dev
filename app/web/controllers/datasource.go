package controllers

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/app/web/form"
	"github.com/kyleu/npn/gen/templates"
	"net/http"
)

func DataSourceList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Breadcrumbs = dsnBreadcrumbs(ctx)
		avail, err := ctx.App.Parsers.Detect(".")
		if err != nil {
			return act.EResp(err, "unable to detect schemata")
		}
		return act.T(templates.DataSourceList(avail, ctx, w))
	})
}

func DataSourceDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		origin := schema.OriginFromString(mux.Vars(r)["t"])
		key := r.URL.Query().Get(util.KeyKey)
		ctx.Breadcrumbs = dsnBreadcrumbs(ctx, "", util.FilenameOf(key))
		sch, rsp, err := ctx.App.Parsers.Load(origin.Key, []string{key})
		if err != nil {
			return act.EResp(err, "unable to calculate schema")
		}
		return act.T(templates.DataSourceDetail(sch, rsp, origin, ctx, w))
	})
}

func DataSourceSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		origin := schema.OriginFromString(mux.Vars(r)["t"])
		frm := &form.SchemaSaveForm{}
		err := form.Decode(r, frm, ctx.Logger)
		if err != nil {
			return act.EResp(err, "invalid form")
		}
		sch, _, err := ctx.App.Parsers.Load(origin.Key, []string{frm.Original})
		if err != nil {
			return act.EResp(err, "unable to calculate schema")
		}
		sch.Key = util.Slugify(frm.Title)
		sch.Title = frm.Title
		err = ctx.App.Files.SaveSchema(sch, true)
		if err != nil {
			return act.EResp(err, "unable to save schema")
		}
		msg := "Schema saved"
		redir := ctx.Route(util.KeySchema + ".detail", util.KeyKey, sch.Key)
		return act.FlashAndRedir(true, msg, redir, w, r, ctx)
	})
}

func dsnBreadcrumbs(ctx *web.RequestContext, pairs ...string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.KeyDataSource), util.KeyDataSource)
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, web.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}
