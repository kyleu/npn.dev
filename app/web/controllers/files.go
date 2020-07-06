package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
)

func FileRoot(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.PluralTitle(util.KeyFile)
		ctx.Breadcrumbs = web.BreadcrumbsSimple("", util.KeyFile)
		files, err := ioutil.ReadDir(".")
		if err != nil {
			return act.EResp(err, "cannot read path [.]")
		}
		return act.T(templates.FileBrowse([]string{}, files, ctx, w))
	})
}

func FilePath(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		path := r.URL.Path
		path = strings.TrimPrefix(path, "/"+util.KeyFile)
		path = strings.TrimPrefix(path, "/")
		paths := strings.Split(path, "/")
		fi, err := os.Stat("./" + path)
		if err != nil {
			return act.EResp(err, "cannot load file ["+path+"]")
		}
		ctx.Title = util.PluralTitle(util.KeyFile)
		ctx.Breadcrumbs = fileBreadcrumbs(ctx, paths...)

		if fi.Mode().IsDir() {
			files, err := ioutil.ReadDir("./" + path)
			if err != nil {
				return act.EResp(err, "cannot read directory ["+path+"]")
			}
			return act.T(templates.FileBrowse(paths, files, ctx, w))
		}
		content, err := ioutil.ReadFile("./" + path)
		if err != nil {
			return act.EResp(err, "cannot read file ["+path+"]")
		}
		return act.T(templates.FileContent(paths, string(content), ctx, w))
	})
}

func fileBreadcrumbs(ctx *web.RequestContext, paths ...string) web.Breadcrumbs {
	route := ctx.Route(util.KeyFile + ".root")
	bc := web.BreadcrumbsSimple(route, util.KeyFile)
	for pathIdx, path := range paths {
		bc = append(bc, web.BreadcrumbsSimple(route+strings.Join(paths[0:pathIdx+1], "/"), path)...)
	}
	return bc
}
