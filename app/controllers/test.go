package controllers

import (
	"net/http"

	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
	"github.com/mccutchen/go-httpbin/httpbin"
)

var testSvc *httpbin.HTTPBin

func TestIndex(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if testSvc == nil {
			testSvc = httpbin.NewHTTPBin()
		}
		ctx.Title = "tests"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("tests")}
		return npncontroller.T(templates.TestList(ctx, w))
	})
}

func TestCall(w http.ResponseWriter, r *http.Request) {
	if testSvc == nil {
		testSvc = httpbin.NewHTTPBin()
	}
	testSvc.Handler().ServeHTTP(w, r)
}
