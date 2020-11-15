package socket

import (
	"encoding/json"

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

func onGetRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &getRequestOut{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load getRequest param")
	}
	req, err := svc.Request.LoadRequest(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestNotFound, frm)
		return s.WriteMessage(c.ID, msg)
	}
	ret := &reqDetailIn{Coll: frm.Coll, Req: req}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onSaveRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &saveRequestOut{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = svc.Request.SaveRequest(&c.Profile.UserID, frm.Coll, frm.Orig, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't save request")
	}
	ret := &reqDetailIn{Coll: frm.Coll, Req: frm.Req}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onDeleteRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &deleteRequestOut{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = svc.Request.DeleteRequest(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't remove request")
	}

	summaries, err := svc.Request.ListRequests(&c.Profile.UserID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't list requests")
	}

	ret := &reqDeleted{Coll: frm.Coll, Req: frm.Req, Requests: summaries}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDeleted, ret)
	return s.WriteMessage(c.ID, msg)
}

func onCall(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := getContext(s)
	frm := &callOut{}
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

func getContext(s *npnconnection.Service) *services {
	return s.Context.(*services)
}
