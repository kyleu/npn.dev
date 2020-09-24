package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

type requestCallParam struct {
	Coll string    `json:"coll"`
	Req  string    `json:"req"`
}

func handleRequestMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case "requestCall":
		frm := &requestCallParam{}
		err := npncore.FromJSONStrict(param, frm)
		if err != nil {
			return errors.Wrap(err, "can't load request param")
		}

		svc := s.Context.(*services)
		req, err := svc.Collection.LoadRequest(frm.Coll, frm.Req)
		if err != nil {
			return errors.Wrap(err, "can't load request ["+frm.Req+"]")
		}

		rsp := svc.Caller.Call(frm.Coll, frm.Req, req.Prototype)

		msg := npnconnection.NewMessage(npncore.KeyRequest, "callResult", rsp)
		err = s.WriteMessage(c.ID, msg)
		return err
	default:
		return errors.New("invalid request command [" + cmd + "]")
	}
}
