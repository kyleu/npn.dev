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

func SchemaList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		avail, err := ctx.App.Parsers.Detect(".")
		schemata := ctx.App.Files.ListSchemata()
		if err != nil {
			return act.EResp(err, "unable to detect schemata")
		}
		return act.T(templates.SchemaList(schemata, avail, ctx, w))
	})
}

func SchemaForm(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		origin := schema.OriginFromString(mux.Vars(r)["t"])
		key := r.URL.Query().Get(util.KeyKey)
		sch, rsp, err := ctx.App.Parsers.Load(origin.Key, key)
		if err != nil {
			return act.EResp(err, "unable to calculate schema")
		}
		return act.T(templates.SchemaSaveForm(sch, rsp, origin, ctx, w))
	})
}

func SchemaSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		origin := schema.OriginFromString(mux.Vars(r)["t"])
		frm := &form.SchemaSaveForm{}
		err := form.Decode(r, frm, ctx.Logger)
		if err != nil {
			return act.EResp(err, "invalid form")
		}
		sch, _, err := ctx.App.Parsers.Load(origin.Key, frm.Original)
		if err != nil {
			return act.EResp(err, "unable to calculate schema")
		}
		sch.Key = util.Slugify(frm.Title)
		sch.Title = frm.Title
		err = ctx.App.Files.SaveSchema(sch, true)
		if err != nil {
			return act.EResp(err, "unablle to save schema")
		}
		msg := "Schema saved"
		redir := ctx.Route(util.KeySchema + ".view", util.KeyKey, sch.Key)
		return act.FlashAndRedir(true, msg, redir, w, r, ctx)
	})
}

func SchemaDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sch, err := ctx.App.Files.LoadSchema(key)
		if err != nil {
			return act.EResp(err, "cannot load schema")
		}
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
