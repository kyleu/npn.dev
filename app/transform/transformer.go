package transform

import "github.com/kyleu/npn/app/request"

type Result struct {
	Key string `json:"key"`
	Out string `json:"out,omitempty"`
}

type Transformer interface {
	Key() string
	Transform(p *request.Prototype) (*Result, error)
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
