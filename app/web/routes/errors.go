package routes

import (
	"context"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"io"
	"net/http"

	"github.com/kyleu/npn/gen/components"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"

	"github.com/kyleu/npn/app/config"

	"github.com/gorilla/mux"
)

func addContext(router *mux.Router, info *config.AppInfo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer internalServerError(router, info, w, r)
		ctx := context.WithValue(r.Context(), util.RoutesKey, router)
		ctx = context.WithValue(ctx, util.InfoKey, info)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func internalServerError(router *mux.Router, info *config.AppInfo, w http.ResponseWriter, r *http.Request) {
	defer lastChanceError(w)

	err := recover()
	if err != nil {
		st := http.StatusInternalServerError

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(st)

		rc := context.WithValue(r.Context(), util.RoutesKey, router)
		rc = context.WithValue(rc, util.InfoKey, info)
		ctx := web.ExtractContext(w, r.WithContext(rc))

		ctx.Title = "Server Error"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(r.URL.Path, util.KeyError)

		e, ok := err.(error)
		if !ok {
			e = errors.New(fmt.Sprintf("err [%v] is of type [%T]", err, err))
		}

		_, _ = components.InternalServerError(npncore.GetErrorDetail(e), r, ctx, w)
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
