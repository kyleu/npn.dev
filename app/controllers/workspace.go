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
		return npncontroller.T(templates.Workspace(ctx, w))
	})
}

func WorkspaceVue(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		return npncontroller.T(templates.WorkspaceVue(ctx, w))
	})
}

var upgrader = websocket.Upgrader{
	HandshakeTimeout:  0,
	ReadBufferSize:    0,
	WriteBufferSize:   0,
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := npnweb.ExtractContext(w, r, true)
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
