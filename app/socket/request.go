package socket

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func handleRequestMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	var err error

	switch cmd {
	case ClientMessageGetRequest:
		err = onGetRequest(c.ID, param, s)
	case ClientMessageSaveRequest:
		err = onSaveRequest(c.ID, param, s)
	case ClientMessageCall:
		err = onCall(c.ID, param, s)
	case ClientMessageTransform:
		err = onTransform(c.ID, param, s)
	default:
		err = errors.New("invalid request command [" + cmd + "]")
	}

	return err
}

func onGetRequest(connID uuid.UUID, param json.RawMessage, s *npnconnection.Service) error {
	svc := s.Context.(*services)
	frm := &paramGetRequest{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load getRequest param")
	}
	req, err := svc.Collection.LoadRequest(frm.Coll, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't load request")
	}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, req)
	return s.WriteMessage(connID, msg)
}

func onSaveRequest(connID uuid.UUID, param json.RawMessage, s *npnconnection.Service) error {
	svc := s.Context.(*services)
	frm := &paramSaveRequest{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = svc.Collection.SaveRequest(frm.Coll, frm.Orig, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't load original request")
	}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestDetail, frm.Req)
	return s.WriteMessage(connID, msg)
}

func onCall(connID uuid.UUID, param json.RawMessage, s *npnconnection.Service) error {
	svc := s.Context.(*services)
	frm := &paramCall{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request call param")
	}

	go func() {
		rsp := svc.Caller.Call(frm.Coll, frm.Req, frm.Proto)
		msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageCallResult, rsp)
		_ = s.WriteMessage(connID, msg)
	}()

	return nil
}

func onTransform(connID uuid.UUID, param json.RawMessage, s *npnconnection.Service) error {
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
	return s.WriteMessage(connID, msg)
}
