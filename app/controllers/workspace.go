package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/gen/templates"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
)

func allow(secret string, r *http.Request) bool {
	if len(secret) > 0 {
		c, _ := r.Cookie("secret")
		if c == nil || c.Value != secret {
			return false
		}
	}
	return true
}

func Enable(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		http.SetCookie(w, &http.Cookie{Name: "secret", Value: ctx.App.Secret()})
		return ctx.Route("workspace"), nil
	})
}

func WorkspaceIndex(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if !allow(ctx.App.Secret(), r) {
			return npncontroller.T(templates.ComingSoon(ctx, w))
		}

		return npncontroller.T(templates.WorkspaceUI(ctx.App.Public(), true, ctx, w))
	})
}

func Workspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if !allow(ctx.App.Secret(), r) {
			return npncontroller.T(templates.ComingSoon(ctx, w))
		}
		return npncontroller.T(templates.WorkspaceUI(ctx.App.Public(), false, ctx, w))
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := npnweb.ExtractContext(w, r, true)

	if !allow(ctx.App.Secret(), r) {
		ctx.Logger.Warn("socket request while locked")
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.Logger.Info(fmt.Sprintf("unable to upgrade connection to websocket: %+v", err))
		return
	}

	ss := app.Svc(ctx.App).Socket

	connID, err := ss.Register(ctx.Profile.ToProfile(), c)
	if err != nil {
		ctx.Logger.Warn("unable to register websocket connection")
		return
	}

	err = ss.OnOpen(connID)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("error processing socket open event: %+v", err))
		return
	}

	err = ss.ReadLoop(connID)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("error processing socket read loop: %+v", err))
		return
	}
}
