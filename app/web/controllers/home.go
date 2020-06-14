package controllers

import (
	"net/http"

	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.AppName
		return act.T(templates.Index(ctx, w))
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.Breadcrumbs{web.BreadcrumbSelf(util.KeyAbout)}
		return act.T(templates.StaticAbout(ctx, w))
	})
}
