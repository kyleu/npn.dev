package transform

import (
	"fmt"
	"github.com/kyleu/npn/app/request"
	"strings"
)

type HTTP struct {
}

func (c *HTTP) Key() string {
	return "http"
}

func (c *HTTP) Transform(p *request.Prototype) (*Result, error) {
	out := []string{}

	var app = func(s string) {
		out = append(out, s)
	}

	app(fmt.Sprintf("%v %v HTTP/1.1", p.Method.Key, p.FullPathString()))
	for _, h := range p.FinalHeaders() {
		app(fmt.Sprintf("%v: %v", h.Key, h.Value))
	}
	return &Result{Out: strings.Join(out, "\n")}, nil
}
