package call

import "github.com/kyleu/npn/app/header"

type Result struct {
	Collection     string         `json:"collection,omitempty"`
	Request        string         `json:"request,omitempty"`
	RequestHeaders header.Headers `json:"requestHeaders,omitempty"`
	Status         string         `json:"status,omitempty"`
	Response       *Response      `json:"response,omitempty"`
	Timing         *Timing        `json:"timing,omitempty"`
	Error          string         `json:"error,omitempty"`
}

func NewResult(coll string, req string, status string) *Result {
	return &Result{Collection: coll, Request: req, Status: status}
}

func NewErrorResult(coll string, req string, err string) *Result {
	ret := NewResult(coll, req, "error")
	ret.Error = err
	return ret
}
