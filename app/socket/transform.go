package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func onTransform(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	frm := &transformIn{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request transform param")
	}

	tx := transform.AllTransformers.Get(frm.Fmt)
	if tx == nil {
		return errors.New("can't load transformer [" + frm.Fmt + "]")
	}

	sess, err := getContext(s).Session.Load(&c.Profile.UserID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load request transform session")
	}

	rsp, err := tx.Transform(frm.Proto, sess)
	if err != nil {
		return errors.Wrap(err, "can't load transform request")
	}

	txr := transformOut{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageTransformResult, txr)
	return s.WriteMessage(c.ID, msg)
}