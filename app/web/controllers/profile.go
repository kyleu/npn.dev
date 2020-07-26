package controllers

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app/web/form"

	"github.com/kyleu/npn/app/web/act"

	"github.com/kyleu/npn/app/web"

	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/gen/templates"
)

type ProfileForm struct {
	Theme     string `mapstructure:"theme"`
	LinkColor string `mapstructure:"linkColor"`
	NavColor  string `mapstructure:"navColor"`
	Ref       string `mapstructure:"ref"`
}

func Profile(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.KeyProfile)}
		ref := r.Header.Get("Referer")
		return act.T(templates.Profile(ref, ctx, w))
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		prof := &ProfileForm{}
		err := form.Decode(r, prof, ctx.Logger)
		if err != nil {
			return act.EResp(err)
		}

		ctx.Profile.Theme = npnuser.ThemeFromString(prof.Theme)
		ctx.Profile.NavColor = prof.NavColor
		ctx.Profile.LinkColor = prof.LinkColor

		err = SaveProfile(ctx.App.Files, ctx.Profile, ctx.Logger)
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
		theme := npnuser.ThemeFromString(key)
		ctx.Profile.Theme = theme
		err := SaveProfile(ctx.App.Files, ctx.Profile, ctx.Logger)
		if err != nil {
			return act.EResp(err, "can't save profile")
		}
		return "", nil
	})
}

func SaveProfile(f *npncore.FileLoader, p *npnuser.UserProfile, logger logur.Logger) error {
	return f.WriteFile("profile.json", npncore.ToJSON(p, logger), true)
}
