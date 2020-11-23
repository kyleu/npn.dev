package transform

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/npncore"
)

type JSON struct {
}

var _ Transformer = (*JSON)(nil)

func (c *JSON) Key() string {
	return "json"
}

func (c *JSON) Transform(p *request.Prototype, sess *session.Session) (*Result, error) {
	out := npncore.ToJSON(p, nil)
	return &Result{Out: out}, nil
}
