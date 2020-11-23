package transform

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
)

type Result struct {
	Key string `json:"key"`
	Out string `json:"out,omitempty"`
}

type Transformer interface {
	Key() string
	Transform(p *request.Prototype, sess *session.Session) (*Result, error)
}

type Transformers []Transformer

func (t Transformers) Get(s string) Transformer {
	for _, x := range t {
		if x.Key() == s {
			return x
		}
	}
	return nil
}

var AllTransformers = Transformers{&CURL{}, &HTTP{}, &JSON{}}
