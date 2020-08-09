package npncontroller

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func RoutesFile(app npnweb.AppInfo, r *mux.Router) {
	file := r.Path(routes.Path(npncore.KeyFile)).Subrouter()
	file.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(FileRoot))).Name(routes.Name(npncore.KeyFile, "root"))
	r.PathPrefix("/" + npncore.KeyFile + "/").Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(FilePath))).Name(routes.Name(npncore.KeyFile))
}

func FileRoot(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.PluralTitle(npncore.KeyFile)
		ctx.Breadcrumbs = npnweb.BreadcrumbsSimple("", npncore.KeyFile)
		files, err := ioutil.ReadDir(".")
		if err != nil {
			return EResp(err, "cannot read path [.]")
		}
		return T(npntemplate.FileBrowse([]string{}, files, ctx, w))
	})
}

func FilePath(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		path := r.URL.Path
		path = strings.TrimPrefix(path, "/"+npncore.KeyFile)
		path = strings.TrimPrefix(path, "/")
		paths := strings.Split(path, "/")
		fi, err := os.Stat("./" + path)
		if err != nil {
			return EResp(err, "cannot load file ["+path+"]")
		}
		ctx.Title = npncore.PluralTitle(npncore.KeyFile)
		ctx.Breadcrumbs = fileBreadcrumbs(ctx, paths...)

		if fi.Mode().IsDir() {
			files, err := ioutil.ReadDir("./" + path)
			if err != nil {
				return EResp(err, "cannot read directory ["+path+"]")
			}
			return T(npntemplate.FileBrowse(paths, files, ctx, w))
		}
		content, err := ioutil.ReadFile("./" + path)
		if err != nil {
			return EResp(err, "cannot read file ["+path+"]")
		}
		return T(npntemplate.FileContent(paths, string(content), ctx, w))
	})
}

func fileBreadcrumbs(ctx *npnweb.RequestContext, paths ...string) npnweb.Breadcrumbs {
	route := ctx.Route(npncore.KeyFile + ".root")
	bc := npnweb.BreadcrumbsSimple(route, npncore.KeyFile)
	for pathIdx, path := range paths {
		bc = append(bc, npnweb.BreadcrumbsSimple(route+strings.Join(paths[0:pathIdx+1], "/"), path)...)
	}
	return bc
}
