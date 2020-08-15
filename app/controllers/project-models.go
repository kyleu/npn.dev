package controllers

import (
	"fmt"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/gen/templates"
)

func ProjectModels(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		ctx.Breadcrumbs = projectBreadcrumbs(ctx, ctx.Route("project.detail", npncore.KeyKey, key), key, "", "models")
		p, err := app.Projects(ctx.App).Load(key)
		if err != nil {
			return npncontroller.EResp(err, "cannot load project ["+key+"]")
		}
		schemata, err := app.Schemata(ctx.App).LoadAll(p.SchemaKeys)
		if err != nil {
			return npncontroller.EResp(err)
		}
		return npncontroller.T(templates.ProjectModels(p, schemata, ctx, w))
	})
}

func ProjectModelsSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		p, err := app.Projects(ctx.App).Load(key)
		if err != nil {
			return npncontroller.EResp(err, "cannot load project ["+key+"]")
		}

		schemata, err := app.Schemata(ctx.App).LoadAll(p.SchemaKeys)
		if err != nil {
			return npncontroller.EResp(err)
		}

		newModels := project.ModelRefs{}

		_ = r.ParseForm()

		for _, sch := range schemata {
			for _, m := range sch.Models {
				_, ok := r.Form[m.String()]
				if ok {
					ref := &project.ModelRef{Pkg: m.Pkg, Key: m.Key}
					pkgOverride := r.Form[fmt.Sprintf("%v::pkgoverride", m.String())]
					if len(pkgOverride) > 0 && len(pkgOverride[0]) > 0 {
						override := util.PkgFromString(pkgOverride[0])
						ref.PkgOverride = override
					}
					newModels = append(newModels, ref)
				}
			}
		}

		p.Models = newModels

		err = app.Projects(ctx.App).Save(key, p, true)
		if err != nil {
			return npncontroller.EResp(err, "cannot save project")
		}

		redir := ctx.Route("project.detail", npncore.KeyKey, key)
		return npncontroller.FlashAndRedir(true, "Saved models", redir, w, r, ctx)
	})
}
