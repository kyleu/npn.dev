package transform

import "github.com/kyleu/npn/app/request"

type Result struct {
	Out string `json:"out,omitempty"`
}

type Transformer interface {
	Key() string
	Transform(p *request.Prototype) (*Result, error)
}

var AllTransformers = []Transformer{&CURL{}, &HTTP{}}
