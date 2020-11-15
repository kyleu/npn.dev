package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func onTransform(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	frm := &transformOut{}
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

	txr := transformIn{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageTransformResult, txr)
	return s.WriteMessage(c.ID, msg)
}
