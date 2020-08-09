package controllers

import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/gen/templates"
)

type schemaSaveForm struct {
	Path  string `mapstructure:"path"`
	Key   string `mapstructure:"key"`
	Title string `mapstructure:"title"`
}

func DataSourceList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Breadcrumbs = dsnBreadcrumbs(ctx)
		avail, err := app.Parsers(ctx.App).Detect(".")
		if err != nil {
			return npncontroller.EResp(err, "unable to detect schemata")
		}
		return npncontroller.T(templates.DataSourceList(avail, ctx, w))
	})
}

func DataSourceDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		t := schema.OriginFromString(mux.Vars(r)["t"])
		key := r.URL.Query().Get(npncore.KeyKey)
		ctx.Breadcrumbs = dsnBreadcrumbs(ctx, "", npncore.FilenameOf(key))
		sch, rsp, err := app.Parsers(ctx.App).Load(t, []string{key})
		if err != nil {
			return npncontroller.EResp(err, "unable to calculate schema")
		}
		return npncontroller.T(templates.DataSourceDetail(sch, rsp, t, ctx, w))
	})
}

func DataSourceSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		t := schema.OriginFromString(mux.Vars(r)["t"])
		frm := &schemaSaveForm{}
		err := npnweb.Decode(r, frm, ctx.Logger)
		if err != nil {
			return npncontroller.EResp(err, "invalid form")
		}
		sch, _, err := app.Parsers(ctx.App).Load(t, []string{frm.Path})
		if err != nil {
			return npncontroller.EResp(err, "unable to calculate schema")
		}
		sch.Key = npncore.Slugify(frm.Title)
		sch.Title = frm.Title
		err = app.Schemata(ctx.App).Save(sch, true)
		if err != nil {
			return npncontroller.EResp(err, "unable to save schema")
		}
		msg := "Schema saved"
		redir := ctx.Route("schema"+".detail", npncore.KeyKey, sch.Key)
		return npncontroller.FlashAndRedir(true, msg, redir, w, r, ctx)
	})
}

func dsnBreadcrumbs(ctx *npnweb.RequestContext, pairs ...string) npnweb.Breadcrumbs {
	bc := npnweb.BreadcrumbsSimple(ctx.Route("datasource"), "datasource")
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, npnweb.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}
