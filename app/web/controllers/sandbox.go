package controllers

import (
	"github.com/kyleu/npn/app/model/sandbox"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"github.com/kyleu/npn/gen/templates"
	"net/http"

	"github.com/gorilla/mux"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.PluralTitle(util.KeySandbox)
		ctx.Breadcrumbs = web.Breadcrumbs{web.BreadcrumbSelf(util.KeySandbox)}
		return act.T(templates.SandboxList(sandbox.AllSandboxes, ctx, w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sb := sandbox.FromString(key)
		if sb == nil {
			return "", util.IDError(util.KeySandbox, key)
		}
		content, rsp, err := sb.Resolve(ctx)
		if err != nil {
			return act.EResp(err, "error running sandbox ["+key+"]")
		}

		ctx.Title = sb.Title + " Sandbox"
		bc := web.BreadcrumbsSimple(ctx.Route(util.KeySandbox), util.KeySandbox)
		bc = append(bc, web.BreadcrumbSelf(key))
		ctx.Breadcrumbs = bc

		return act.T(templates.SandboxRun(sb, content, util.ToJSON(rsp, ctx.Logger), ctx, w))
	})
}
