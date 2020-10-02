package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/app/request"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

type requestCallParam struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req"`
	Proto *request.Prototype `json:"proto"`
}

func handleRequestMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	svc := s.Context.(*services)
	var err error

	switch cmd {
	case ClientMessageGetRequest:
		frm := &requestCallParam{}
		err := npncore.FromJSONStrict(param, frm)
		if err != nil {
			return errors.Wrap(err, "can't load request param")
		}
		req, err := svc.Collection.LoadRequest(frm.Coll, frm.Req)
		if err != nil {
			return errors.Wrap(err, "can't load request")
		}
		msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, req)
		err = s.WriteMessage(c.ID, msg)
		if err != nil {
			return errors.Wrap(err, "can't write message")
		}
	case ClientMessageCall:
		frm := &requestCallParam{}
		err := npncore.FromJSONStrict(param, frm)
		if err != nil {
			return errors.Wrap(err, "can't load request param")
		}

		go func() {
			rsp := svc.Caller.Call(frm.Coll, frm.Req, frm.Proto)
			println(rsp.Status)
			msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageCallResult, rsp)
			println(msg.Cmd)
			_ = s.WriteMessage(c.ID, msg)
		}()

	default:
		err = errors.New("invalid request command [" + cmd + "]")
	}

	return err
}
