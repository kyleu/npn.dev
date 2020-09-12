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

func Workspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		title := "Workspace"

		ctx.Title = title
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("workspace")}

		debug := "TODO"

		return npncontroller.T(templates.Workspace(title, debug, ctx, w))
	})
}

var upgrader = websocket.Upgrader{}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := npnweb.ExtractContext(w, r, true)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.Logger.Info("unable to upgrade connection to websocket")
		return
	}

	connID, err := app.Svc(ctx.App).Socket.Register(ctx.Profile.ToProfile(), c)
	if err != nil {
		ctx.Logger.Warn("unable to register websocket connection")
		return
	}

	err = app.Svc(ctx.App).Socket.ReadLoop(connID)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("error processing socket read loop: %+v", err))
		return
	}
}
