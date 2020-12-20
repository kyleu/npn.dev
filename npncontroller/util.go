package npncontroller

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"
)

// Common utility routes, for viewing the installed routes and modules, and providing sitemap.xml, OPTIONS, and 404 handlers
func RoutesUtil(app npnweb.AppInfo, r *mux.Router) {
	rts := r.Path(routes.Path(npncore.KeyRoutes)).Subrouter()
	rts.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RouteList))).Name(routes.Name(npncore.KeyRoutes))

	modules := r.Path(routes.Path(npncore.KeyModules)).Subrouter()
	modules.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ModuleList))).Name(routes.Name(npncore.KeyModules))

	sitemap := r.Path(routes.Path("sitemap.xml")).Subrouter()
	sitemap.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SitemapXML))).Name(routes.Name("sitemap"))

	r.PathPrefix("").Methods(http.MethodOptions).Handler(routes.AddContext(r, app, http.HandlerFunc(Options)))
	r.PathPrefix("").Handler(routes.AddContext(r, app, http.HandlerFunc(NotFound)))
}

// List the configured HTTP routes
func RouteList(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.Title(npncore.KeyRoutes)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyRoutes)}
		return T(npntemplate.RoutesList(ctx, w))
	})
}

// Lists the Go modules used
func ModuleList(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.Title(npncore.KeyModules)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyModules)}
		return T(npntemplate.ModulesList(ctx, w))
	})
}

// Handles the standard /sitemap.xml request for the configured routes
func SitemapXML(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ret := make([]string, 0)
		ret = append(ret, `<?xml version="1.0" encoding="UTF-8"?>`)
		ret = append(ret, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
		for _, rt := range npnweb.ExtractRoutes(ctx.Routes) {
			if routeMatches(rt) {
				url := rt.Path
				ret = append(ret, `  <url>`)
				ret = append(ret, `     <loc>`+url+`</loc>`)
				ret = append(ret, `     <changefreq>always</changefreq>`)
				ret = append(ret, `  </url>`)
			}
		}
		ret = append(ret, `</urlset>`)
		_, _ = w.Write([]byte(strings.Join(ret, "\n")))
		return "", nil
	})
}

func routeMatches(rt *npnweb.RouteDescription) bool {
	pathCheck := func(s ...string) bool {
		for _, x := range s {
			if strings.Contains(rt.Path, x) {
				return false
			}
		}
		return true
	}
	if !pathCheck("admin", "assets", "sitemap", "robots", "{") {
		return false
	}
	if rt.Path == "/ws" {
		return false
	}
	if !strings.Contains(rt.Methods, http.MethodGet) {
		return false
	}
	return true
}
