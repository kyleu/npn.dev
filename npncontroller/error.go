package npncontroller

import (
	"fmt"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	WriteCORS(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	ctx := npnweb.ExtractContext(w, r, false)
	ctx.Title = "Not Found"
	ctx.Breadcrumbs = npnweb.BreadcrumbsSimple(r.URL.Path, "not found")
	ctx.Logger.Info(fmt.Sprintf("[%v %v] returned [%d]", r.Method, r.URL.Path, http.StatusNotFound))
	_, _ = npntemplate.NotFound(r, ctx, w)
}

func Options(w http.ResponseWriter, r *http.Request) {
	WriteCORS(w)
	w.WriteHeader(http.StatusOK)
}
