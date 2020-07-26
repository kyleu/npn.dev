package controllers

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"

	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = npncore.AppName
		return act.T(templates.Index(ctx, w))
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "About " + npncore.AppName
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.KeyAbout)}
		return act.T(templates.StaticAbout(ctx, w))
	})
}

func RouteList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.Title(util.KeyRoutes)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.KeyRoutes)}
		return act.T(templates.RoutesList(ctx, w))
	})
}

func ModuleList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.Title(util.KeyModules)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.KeyModules)}
		return act.T(templates.ModulesList(ctx, w))
	})
}
