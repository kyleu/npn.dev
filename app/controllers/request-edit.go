package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/gen/templates"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
)

func RequestNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll := mux.Vars(r)["c"]
		req := request.NewRequest()
		return npncontroller.T(templates.RequestForm(coll, req, nil, ctx, w))
	})
}

func RequestDelete(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll := mux.Vars(r)["c"]
		key := mux.Vars(r)[npncore.KeyKey]

		err := app.Svc(ctx.App).Collection.DeleteRequest(coll, key)
		if err != nil {
			return npncontroller.EResp(err, "unable to delete ["+coll+"/"+key+"]")
		}

		msg := "deleted request [" + key + "] from this collection"
		return npncontroller.FlashAndRedir(true, msg, ctx.Route(KeyCollection+".detail", "c", coll), w, r, ctx)
	})
}

func RequestEdit(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		coll, req, err := loadRequest(r, ctx, "edit")
		if err != nil {
			return npncontroller.EResp(err)
		}
		return npncontroller.T(templates.RequestForm(coll, req, nil, ctx, w))
	})
}

func RequestSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		frm := &requestForm{}
		err := npnweb.Decode(r, frm, ctx.Logger)
		if err != nil {
			return npncontroller.EResp(err)
		}
		req, err := frm.ToRequest()
		if err != nil {
			return npncontroller.EResp(err, "unable to parse request")
		}

		err = app.Svc(ctx.App).Collection.SaveRequest(frm.Coll, frm.OriginalKey, req)
		if err != nil {
			return npncontroller.EResp(err, "unable to save ["+frm.Coll+"/"+req.Key+"]")
		}

		msg := "saved request [" + req.Key + "]"
		rt := ctx.Route(KeyRequest, "c", frm.Coll, npncore.KeyKey, req.Key)
		return npncontroller.FlashAndRedir(true, msg, rt, w, r, ctx)
	})
}
