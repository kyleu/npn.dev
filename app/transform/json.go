package transform

import (
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"logur.dev/logur"
)

type JSON struct {
}

var _ RequestTransformer = (*JSON)(nil)

func (x *JSON) Key() string {
	return "json"
}

func (x *JSON) Title() string {
	return "JSON"
}

func (x *JSON) Description() string {
	return "TODO: json"
}

func (x *JSON) TransformRequest(proto *request.Prototype, sess *session.Session, logger logur.Logger) (*Result, error) {
	out := npncore.ToJSON(proto, nil)
	return &Result{Out: out}, nil
}

func (x *JSON) TransformCollection(c *collection.FullCollection, logger logur.Logger) (*Result, error) {
	src := map[string]interface{}{"coll": c.Coll, "requests": c.Requests}
	out := npncore.ToJSON(src, nil)
	return &Result{Out: out}, nil
}

var txJSON = &JSON{}
