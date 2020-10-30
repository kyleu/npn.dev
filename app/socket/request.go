package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/app/request"

	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func handleRequestMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	var err error

	switch cmd {
	case ClientMessageGetRequest:
		err = onGetRequest(c, param, s)
	case ClientMessageSaveRequest:
		err = onSaveRequest(c, param, s)
	case ClientMessageDeleteRequest:
		err = onDeleteRequest(c, param, s)
	case ClientMessageCall:
		err = onCall(c, param, s)
	case ClientMessageTransform:
		err = onTransform(c, param, s)
	default:
		err = errors.New("invalid request command [" + cmd + "]")
	}

	return err
}

type reqDetail struct {
	Coll string           `json:"coll"`
	Req  *request.Request `json:"req"`
}

func onGetRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &paramGetRequest{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load getRequest param")
	}
	req, err := svc.Collection.LoadRequest(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't load request")
	}
	ret := &reqDetail{Coll: frm.Coll, Req: req}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onSaveRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &paramSaveRequest{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = svc.Collection.SaveRequest(&c.Profile.UserID, frm.Coll, frm.Orig, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't save request")
	}
	ret := &reqDetail{Coll: frm.Coll, Req: frm.Req}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onDeleteRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &paramDeleteRequest{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = svc.Collection.DeleteRequest(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't remove request")
	}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDeleted, frm.Req)
	return s.WriteMessage(c.ID, msg)
}

func onCall(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &paramCall{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request call param")
	}

	go func() {
		rsp := svc.Caller.Call(frm.Coll, frm.Req, frm.Proto)
		msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageCallResult, rsp)
		_ = s.WriteMessage(c.ID, msg)
	}()

	return nil
}

func onTransform(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	frm := &paramTransform{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request transform param")
	}

	tx := transform.AllTransformers.Get(frm.Fmt)
	if tx == nil {
		return errors.New("can't load transformer [" + frm.Fmt + "]")
	}

	rsp, err := tx.Transform(frm.Proto)
	if err != nil {
		return errors.Wrap(err, "can't load transform request")
	}

	txr := transformResponse{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageTransformResult, txr)
	return s.WriteMessage(c.ID, msg)
}

func getContext(s *npnconnection.Service) *services {
	return s.Context.(*services)
}
