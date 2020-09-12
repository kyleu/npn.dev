package transform

import "github.com/kyleu/npn/app/request"

type Result struct {
	Out string `json:"out,omitempty"`
}

type Transformer interface {
	Key() string
	Transform(r *request.Request) (*Result, error)
}
