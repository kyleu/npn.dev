package transform

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
)

type HTTP struct {
}

var _ RequestTransformer = (*HTTP)(nil)

func (x *HTTP) Key() string {
	return "http"
}

func (x *HTTP) Title() string {
	return "HTTP"
}

func (x *HTTP) Description() string {
	return "TODO: http"
}

func (x *HTTP) ApplyToMultiple() bool {
	return false
}

func (x *HTTP) TransformRequest(p *request.Prototype, sess *session.Session, logger *logrus.Logger) (*Result, error) {
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

var txHTTP = &HTTP{}
