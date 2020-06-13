package controllers

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/web/form"

	"github.com/kyleu/npn/app/web/act"

	"github.com/kyleu/npn/app/web"

	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/gen/templates"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.KeyProfile), util.KeyProfile)
		ref := r.Header.Get("Referer")
		return act.T(templates.Profile(ref, ctx, w))
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		prof := &form.ProfileForm{}
		err := form.Decode(r, prof, ctx.Logger)
		if err != nil {
			return act.EResp(err)
		}

		ctx.Profile.Theme = util.ThemeFromString(prof.Theme)
		ctx.Profile.NavColor = prof.NavColor
		ctx.Profile.LinkColor = prof.LinkColor

		err = ctx.App.Files.SaveProfile(ctx.Profile)
		if err != nil {
			return act.EResp(err, "unable to save profile")
		}
		ref := strings.TrimSpace(prof.Ref)
		if len(ref) == 0 || strings.HasPrefix(ref, "http") {
			ref = "home"
		}
		return act.FlashAndRedir(true, "Profile saved", ref, w, r, ctx)
	})
}

func ProfileTheme(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		theme := util.ThemeFromString(key)
		ctx.Profile.Theme = theme
		err := ctx.App.Files.SaveProfile(ctx.Profile)
		if err != nil {
			return act.EResp(err, "can't save profile")
		}
		return "", nil
	})
}
