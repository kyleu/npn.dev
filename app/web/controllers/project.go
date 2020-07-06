package controllers

import (
	"emperror.dev/errors"
	"net/http"
	"strings"

	"github.com/kyleu/npn/app/model/project"
	"github.com/kyleu/npn/app/web/form"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
)

func ProjectList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Breadcrumbs = projectBreadcrumbs(ctx)
		summaries, err := ctx.App.Projects.Summaries()
		if err != nil {
			return act.EResp(err, "cannot load projects")
		}
		return act.T(templates.ProjectList(summaries, ctx, w))
	})
}

func ProjectNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, "", "new")
		schemaSummaries, err := ctx.App.Schemata.Summaries()
		if err != nil {
			return act.EResp(err, "can't load schema summaries")
		}
		return act.T(templates.ProjectForm(project.DefaultProject, schemaSummaries, ctx, w))
	})
}

func ProjectEdit(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, "", key)
		p, err := ctx.App.Projects.Load(key)
		if err != nil {
			return act.EResp(err, "cannot load project ["+key+"]")
		}
		schemaSummaries, err := ctx.App.Schemata.Summaries()
		if err != nil {
			return act.EResp(err, "can't load schema summaries")
		}
		return act.T(templates.ProjectForm(p, schemaSummaries, ctx, w))
	})
}

func ProjectSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		originalKey := mux.Vars(r)[util.KeyKey]
		frm := &form.ProjectSaveForm{}
		err := form.Decode(r, frm, ctx.Logger)
		if err != nil {
			return act.EResp(err, "invalid project form")
		}
		newKey := util.Slugify(frm.Title)
		if len(newKey) == 0 {
			return act.EResp(errors.New("title is required"))
		}
		pkg := strings.Split(frm.Pkg, "||")
		proj := &project.Project{Key: newKey, Title: frm.Title, RootPath: frm.Path, RootPkg: pkg, SchemaKeys: strings.Split(frm.Schema, "||")}
		err = ctx.App.Projects.Save(originalKey, proj, true)
		if err != nil {
			return act.EResp(err, "cannot save project")
		}

		redir := ctx.Route(util.KeyProject+".detail", util.KeyKey, newKey)
		return act.FlashAndRedir(true, "Saved project", redir, w, r, ctx)
	})
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, "", key)
		p, err := ctx.App.Projects.Load(key)
		if err != nil {
			return act.EResp(err, "cannot load project ["+key+"]")
		}
		return act.T(templates.ProjectDetail(p, ctx, w))
	})
}

func projectBreadcrumbs(ctx *web.RequestContext, pairs ...string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.KeyProject), util.KeyProject)
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, web.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}
