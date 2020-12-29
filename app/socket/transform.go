package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/app/collection"

	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
)

func onTransformRequest(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	frm := &transformIn{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request transform param")
	}

	tx := transform.AllRequestTransformers.Get(frm.Fmt)
	if tx == nil {
		return errors.New("can't load request transformer [" + frm.Fmt + "]")
	}

	sess, err := ctx(s).Session.Load(&c.Profile.UserID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load request transform session ["+frm.Sess+"]")
	}

	rsp, err := tx.TransformRequest(frm.Proto, sess, s.Logger)
	if err != nil {
		return errors.Wrap(err, "can't transform request")
	}

	txr := transformOut{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestTransform, txr)
	return s.WriteMessage(c.ID, msg)
}

func onTransformCollection(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	frm := &transformIn{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform param")
	}

	tx := transform.AllCollectionTransformers.Get(frm.Fmt)
	if tx == nil {
		return errors.New("can't load collection transformer [" + frm.Fmt + "]")
	}

	svcs := ctx(s)

	coll, err := svcs.Collection.Load(&c.Profile.UserID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform collection ["+frm.Coll+"]")
	}

	requests, err := svcs.Request.LoadAll(&c.Profile.UserID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform requests for ["+frm.Coll+"]")
	}

	sess, err := ctx(s).Session.Load(&c.Profile.UserID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform session ["+frm.Sess+"]")
	}

	tc := &collection.FullCollection{Coll: coll, Requests: requests, Sess: sess}
	rsp, err := tx.TransformCollection(tc, s.Logger)
	if err != nil {
		return errors.Wrap(err, "can't load transform colllection")
	}

	txr := transformOut{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionTransform, txr)
	return s.WriteMessage(c.ID, msg)
}
