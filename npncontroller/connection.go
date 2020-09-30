package npncontroller

import (
	"encoding/json"
	"net/http"

	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/gorilla/mux"
)

type connectionForm struct {
	ID    string `mapstructure:"id"`
	Svc   string `mapstructure:"svc"`
	Cmd   string `mapstructure:"cmd"`
	Param string `mapstructure:"param"`
}

var globalSvc *npnconnection.Service

func RoutesSocketAdmin(app npnweb.AppInfo, svc *npnconnection.Service, r *mux.Router) {
	globalSvc = svc
	sr := r.Path(routes.Adm(npncore.KeyConnection)).Subrouter()
	sr.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ConnectionList))).Name(npnweb.AdminLink(npncore.KeyConnection))
	sr.Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(ConnectionPost))).Name(npnweb.AdminLink(npncore.KeyConnection, "post"))
	r.Path(routes.Adm(npncore.KeyConnection, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ConnectionDetail))).Name(npnweb.AdminLink(npncore.KeyConnection, npncore.KeyDetail))
}

func ConnectionList(w http.ResponseWriter, r *http.Request) {
	AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Connection List"
		ctx.Breadcrumbs = AdminBC(ctx, npncore.KeyConnection, npncore.Plural(npncore.KeyConnection))

		p := npnweb.ParamSetFromRequest(r)
		connections := globalSvc.List(p.Get(npncore.KeySocket, ctx.Logger))
		return T(npntemplate.ConnectionList(connections, p, ctx, w))
	})
}

func ConnectionDetail(w http.ResponseWriter, r *http.Request) {
	AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		connectionID, err := npnweb.IDFromParams(npncore.KeyConnection, mux.Vars(r))
		if err != nil {
			return EResp(err)
		}
		connection := globalSvc.GetByID(*connectionID)
		ctx.Title = connection.ID.String()
		bc := AdminBC(ctx, npncore.KeyConnection, npncore.Plural(npncore.KeyConnection))
		str := connectionID.String()
		bc = append(bc, npnweb.BreadcrumbSelf(str[0:8]))
		ctx.Breadcrumbs = bc

		msg := npnconnection.NewMessage(npncore.KeySystem, npncore.KeyTest, nil)
		return T(npntemplate.ConnectionDetail(connection, msg, ctx, w))
	})
}

func ConnectionPost(w http.ResponseWriter, r *http.Request) {
	AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		frm := &connectionForm{}
		err := npnweb.Decode(r, frm, ctx.Logger)
		if err != nil {
			return EResp(err)
		}

		connectionID := npncore.GetUUIDFromString(frm.ID)
		connection := globalSvc.GetByID(*connectionID)

		var param []map[string]interface{}
		_ = json.Unmarshal([]byte(frm.Param), &param)
		msg := npnconnection.NewMessage(frm.Svc, frm.Cmd, param)
		err = globalSvc.WriteMessage(*connectionID, msg)
		if err != nil {
			return EResp(err)
		}

		ctx.Title = connectionID.String()
		bc := AdminBC(ctx, npncore.KeyConnection, npncore.Plural(npncore.KeyConnection))
		str := connectionID.String()
		bc = append(bc, npnweb.BreadcrumbSelf(str[0:8]))
		ctx.Breadcrumbs = bc

		return T(npntemplate.ConnectionDetail(connection, msg, ctx, w))
	})
}
