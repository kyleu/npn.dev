package transform

import (
	"fmt"
	"github.com/kyleu/npn/app/session"
	"strings"

	"github.com/kyleu/npn/app/request"
)

type CURL struct {
	Silent    bool
	Multiline bool
}

var _ Transformer = (*CURL)(nil)

func (c *CURL) Key() string {
	return "curl"
}

func (c *CURL) Transform(p *request.Prototype, sess *session.Session) (*Result, error) {
	out := []string{"curl"}

	var app = func(s string) {
		out = append(out, s)
	}
	var esc = func(s string) string {
		return strings.ReplaceAll(s, "'", "'\\''")
	}

	if c.Silent {
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
	if !c.Multiline {
		sep = " \\\n  "
	}
	return &Result{Out: strings.Join(out, sep)}, nil
}
