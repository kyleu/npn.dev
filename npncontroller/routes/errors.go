package routes

import (
	"context"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"
	"io"
	"net/http"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
)

func AddContext(router *mux.Router, info npnweb.AppInfo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer internalServerError(router, info, w, r)
		rt := context.WithValue(r.Context(), npncore.RoutesKey, router)
		ctx := context.WithValue(rt, npncore.InfoKey, info)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func internalServerError(router *mux.Router, info npnweb.AppInfo, w http.ResponseWriter, r *http.Request) {
	defer lastChanceError(w)

	err := recover()
	if err != nil {
		st := http.StatusInternalServerError

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(st)

		rc := context.WithValue(r.Context(), npncore.RoutesKey, router)
		rc = context.WithValue(rc, npncore.InfoKey, info)
		ctx := npnweb.ExtractContext(w, r.WithContext(rc), false)

		ctx.Title = "Server Error"
		ctx.Breadcrumbs = npnweb.BreadcrumbsSimple(r.URL.Path, npncore.KeyError)

		e, ok := err.(error)
		if !ok {
			e = errors.New(fmt.Sprintf("err [%v] is of type [%T]", err, err))
		}

		_, _ = npntemplate.InternalServerError(npncore.GetErrorDetail(e), r, ctx, w)
		ctx.Logger.Warn(fmt.Sprintf("[%v %v] returned [%d]: %+v", r.Method, r.URL.Path, st, e))
	}
}

func lastChanceError(w io.Writer) {
	err := recover()
	if err != nil {
		fmt.Println(fmt.Sprintf("unhandled last-chance error while processing error handler: %+v", err))
		_, _ = w.Write([]byte("Internal Server Error"))
	}
}
