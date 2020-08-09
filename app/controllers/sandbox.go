package controllers

import (
	"fmt"
	"github.com/kyleu/npn/app/sandbox"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"

	"github.com/kyleu/npn/gen/templates"

	"github.com/gorilla/mux"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.PluralTitle(npncore.KeySandbox)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeySandbox)}
		return npncontroller.T(templates.SandboxList(sandbox.All(), ctx, w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sb := sandbox.FromString(key)
		if sb == nil {
			return "", npncore.IDError(npncore.KeySandbox, key)
		}
		content, rsp, err := sb.Resolve(ctx)
		if err != nil {
			return npncontroller.EResp(err, "error running sandbox ["+key+"]")
		}

		ctx.Title = sb.Title + " Sandbox"
		bc := npnweb.BreadcrumbsSimple(ctx.Route(npncore.KeySandbox), npncore.KeySandbox)
		bc = append(bc, npnweb.BreadcrumbSelf(key))
		ctx.Breadcrumbs = bc
		t := fmt.Sprintf("%T", rsp)

		return npncontroller.T(templates.SandboxRun(sb, t, content, npncore.ToJSON(rsp, ctx.Logger), ctx, w))
	})
}
