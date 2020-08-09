package controllers

import (
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/npn/app/project"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/gen/templates"
)

type projectSaveForm struct {
	Title     string `mapstructure:"title"`
	Org       string `mapstructure:"org"`
	Schema    string `mapstructure:"schema"`
	Path      string `mapstructure:"path"`
	Pkg       string `mapstructure:"pkg"`
	Prototype string `mapstructure:"proto"`
	Options   string `mapstructure:"options"`
}

func ProjectList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Breadcrumbs = projectBreadcrumbs(ctx)
		summaries, err := app.Projects(ctx.App).Summaries()
		if err != nil {
			return npncontroller.EResp(err, "cannot load projects")
		}
		return npncontroller.T(templates.ProjectList(summaries, ctx, w))
	})
}

func ProjectNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, "", "new")
		schemaSummaries, err := app.Schemata(ctx.App).Summaries()
		if err != nil {
			return npncontroller.EResp(err, "can't load schema summaries")
		}
		return npncontroller.T(templates.ProjectForm(project.DefaultProject, schemaSummaries, ctx, w))
	})
}

func ProjectEdit(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, "", key)
		p, err := app.Projects(ctx.App).Load(key)
		if err != nil {
			return npncontroller.EResp(err, "cannot load project ["+key+"]")
		}
		schemaSummaries, err := app.Schemata(ctx.App).Summaries()
		if err != nil {
			return npncontroller.EResp(err, "can't load schema summaries")
		}
		return npncontroller.T(templates.ProjectForm(p, schemaSummaries, ctx, w))
	})
}

func ProjectSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		originalKey := mux.Vars(r)[npncore.KeyKey]
		frm := &projectSaveForm{}
		err := npnweb.Decode(r, frm, ctx.Logger)
		if err != nil {
			return npncontroller.EResp(err, "invalid project form")
		}
		newKey := npncore.Slugify(frm.Title)
		if len(newKey) == 0 {
			return npncontroller.EResp(errors.New("title is required"))
		}

		opts := map[string]interface{}{}
		for _, v := range strings.Split(frm.Options, "||") {
			v = strings.TrimSpace(v)
			if len(v) > 0 {
				idx := strings.Index(v, "::")
				if idx == -1 {
					opts[v] = v
				} else {
					opts[v[0:idx]] = v[idx+1:]
				}
			}
		}

		org := frm.Org
		if len(org) == 0 {
			org = newKey
		}
		proj := &project.Project{
			Key:        newKey,
			Title:      frm.Title,
			Org:        org,
			RootPath:   frm.Path,
			RootPkg:    strings.Split(frm.Pkg, "||"),
			Prototype:  frm.Prototype,
			SchemaKeys: strings.Split(frm.Schema, "||"),
			Options:    opts,
		}
		err = app.Projects(ctx.App).Save(originalKey, proj, true)
		if err != nil {
			return npncontroller.EResp(err, "cannot save project")
		}

		redir := ctx.Route("project.detail", npncore.KeyKey, newKey)
		return npncontroller.FlashAndRedir(true, "Saved project", redir, w, r, ctx)
	})
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, "", key)
		p, err := app.Projects(ctx.App).Load(key)
		if err != nil {
			return npncontroller.EResp(err, "cannot load project ["+key+"]")
		}
		return npncontroller.T(templates.ProjectDetail(p, ctx, w))
	})
}

func projectBreadcrumbs(ctx *npnweb.RequestContext, pairs ...string) npnweb.Breadcrumbs {
	bc := npnweb.BreadcrumbsSimple(ctx.Route("project"), "project")
	for i := 0; i < len(pairs)-1; i += 2 {
		bc = append(bc, npnweb.BreadcrumbsSimple(pairs[i], pairs[i+1])...)
	}
	return bc
}
