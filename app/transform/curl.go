package transform

import "github.com/kyleu/npn/app/request"

type CURL struct {

}

func (c *CURL) Key() string {
	return "curl"
}

func (c *CURL) Transform(r *request.Request) (*Result, error) {
	out := "TODO"
	return &Result{Out: out}, nil
}
