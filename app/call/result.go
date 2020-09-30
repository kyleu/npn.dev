package call

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/npn/npncore"
)

type Result struct {
	ID             uuid.UUID      `json:"id,omitempty"`
	URL            string         `json:"url,omitempty"`
	Collection     string         `json:"collection,omitempty"`
	Request        string         `json:"request,omitempty"`
	RequestHeaders header.Headers `json:"requestHeaders,omitempty"`
	Status         string         `json:"status,omitempty"`
	RedirectedFrom *Result        `json:"redirectedFrom,omitempty"`
	Response       *Response      `json:"response,omitempty"`
	Timing         *Timing        `json:"timing,omitempty"`
	Error          string         `json:"error,omitempty"`
}

func NewResult(url string, coll string, req string, status string) *Result {
	return &Result{ID: npncore.UUID(), URL: url, Collection: coll, Request: req, Status: status}
}

func NewErrorResult(url string, coll string, req string, err string) *Result {
	ret := NewResult(url, coll, req, "error")
	ret.Error = err
	return ret
}
