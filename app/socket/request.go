package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/app/call"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func handleRequestMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	var err error

	switch cmd {
	case ClientMessageRunURL:
		err = onRunURL(c, param, s)
	case ClientMessageGetRequest:
		err = onGetRequest(c, param, s)
	case ClientMessageSaveRequest:
		err = onSaveRequest(c, param, s)
	case ClientMessageDeleteRequest:
		err = onDeleteRequest(c, param, s)
	case ClientMessageCall:
		err = onCall(c, param, s)
	case ClientMessageTransform:
		err = onTransformRequest(c, param, s)
	default:
		err = errors.New("invalid request command [" + cmd + "]")
	}

	return err
}

func onRunURL(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	url, err := npncore.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to read URL")
	}
	println("TODO: onRunURL(" + url + ")")
	return nil
}

func onGetRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := ctx(s)
	frm := &getRequestIn{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load getRequest param")
	}
	req, err := svc.Request.LoadRequest(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestNotFound, frm)
		return s.WriteMessage(c.ID, msg)
	}
	ret := &reqDetailOut{Coll: frm.Coll, Req: req}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onSaveRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := ctx(s)
	frm := &saveRequestIn{}
	err := npncore.FromJSON(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	frm.Req = frm.Req.Minify()
	err = svc.Request.SaveRequest(&c.Profile.UserID, frm.Coll, frm.Orig, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't save request ["+frm.Req.Key+"]")
	}
	ret := &reqDetailOut{Coll: frm.Coll, Req: frm.Req}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onDeleteRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := ctx(s)
	frm := &deleteRequestIn{}
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

	ret := &reqDeletedOut{Coll: frm.Coll, Req: frm.Req, Requests: summaries}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDeleted, ret)
	return s.WriteMessage(c.ID, msg)
}

func onCall(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := ctx(s)
	frm := &callIn{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request call param")
	}

	sess, err := ctx(s).Session.Load(&c.Profile.UserID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load session ["+frm.Sess+"]")
	}

	go func() {
		onStarted := func(started *call.RequestStarted) {
			msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestStarted, started)
			_ = s.WriteMessage(c.ID, msg)
		}
		onCompleted := func(completed *call.RequestCompleted) {
			msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestCompleted, completed)
			_ = s.WriteMessage(c.ID, msg)
		}
		_ = svc.Caller.Call(&c.Profile.UserID, frm.Coll, frm.Req, frm.Proto, sess, onStarted, onCompleted)
	}()

	return nil
}
