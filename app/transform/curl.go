package transform

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
)

type CURL struct {
	Silent    bool
	Multiline bool
}

var _ RequestTransformer = (*CURL)(nil)

func (x *CURL) Key() string {
	return "curl"
}

func (x *CURL) Title() string {
	return "CURL"
}

func (x *CURL) Description() string {
	return "TODO: curl"
}

func (x *CURL) ApplyToMultiple() bool {
	return false
}

func (x *CURL) TransformRequest(p *request.Prototype, sess *session.Session, logger *logrus.Logger) (*Result, error) {
	out := []string{"curl"}

	var app = func(s string) {
		out = append(out, s)
	}
	var esc = func(s string) string {
		return strings.ReplaceAll(s, "'", "'\\''")
	}

	if x.Silent {
		app("--silent")
	}
	if p.Options != nil && p.Options.Timeout > 0 {
		app(fmt.Sprintf("--max-time %v", p.Options.Timeout))
	}
	if p.Options != nil && (!p.Options.IgnoreRedirects) {
		app("--location")
	}
	if p.Method == request.MethodHead {
		app("--head")
	} else {
		app("--request " + p.Method.Key)
	}

	for _, h := range p.Headers {
		app(fmt.Sprintf("--header '%v: %v'", esc(h.Key), esc(h.Value)))
	}

	app("'" + esc(p.URLString()) + "'")

	sep := " "
	if !x.Multiline {
		sep = " \\\n  "
	}
	return &Result{Out: strings.Join(out, sep)}, nil
}

var txCURL = &CURL{}
