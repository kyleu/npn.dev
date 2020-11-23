package transform

import (
	"fmt"
	"github.com/kyleu/npn/app/session"
	"strings"

	"github.com/kyleu/npn/app/request"
)

type HTTP struct {
}

var _ Transformer = (*HTTP)(nil)

func (c *HTTP) Key() string {
	return "http"
}

func (c *HTTP) Transform(p *request.Prototype, sess *session.Session) (*Result, error) {
	out := []string{}

	var app = func(s string) {
		out = append(out, s)
	}

	app(fmt.Sprintf("%v %v HTTP/1.1", p.Method.Key, p.FullPathString()))
	for _, h := range p.FinalHeaders(sess) {
		app(fmt.Sprintf("%v: %v", h.Key, h.Value))
	}
	return &Result{Out: strings.Join(out, "\n")}, nil
}
