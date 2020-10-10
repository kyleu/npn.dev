package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/npn/gen/templates"
)

func System(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.AppName
		return npncontroller.T(templates.SystemIndex(ctx, w))
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "About " + npncore.AppName
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyAbout)}
		return npncontroller.T(templates.StaticAbout(ctx, w))
	})
}

func Source(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		http.ServeFile(w, r, "./client/src/"+r.URL.Path)
		return "", nil
	})
}
