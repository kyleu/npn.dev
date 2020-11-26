package transform

import (
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type JSON struct {
}

var _ RequestTransformer = (*JSON)(nil)

func (x *JSON) Key() string {
	return "json"
}

func (x *JSON) Description() string {
	return "TODO: json"
}

func (x *JSON) TransformRequest(proto *request.Prototype, sess *session.Session, logger logur.Logger) (*Result, error) {
	out := npncore.ToJSON(proto, nil)
	return &Result{Out: out}, nil
}

func (x *JSON) TransformCollection(coll *collection.Collection, requests request.Requests, sess *session.Session, logger logur.Logger) (*Result, error) {
	src := map[string]interface{}{"coll": coll, "requests": requests}
	out := npncore.ToJSON(src, nil)
	return &Result{Out: out}, nil
}

var txJSON = &JSON{}
