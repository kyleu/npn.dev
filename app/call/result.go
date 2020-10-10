package call

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
)

type Result struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Collection string    `json:"collection,omitempty"`
	Request    string    `json:"request,omitempty"`
	Status     string    `json:"status,omitempty"`
	Response   *Response `json:"response,omitempty"`
	Error      string    `json:"error,omitempty"`
}

func NewResult(coll string, req string, status string) *Result {
	return &Result{ID: npncore.UUID(), Collection: coll, Request: req, Status: status}
}

func NewErrorResult(coll string, req string, err string) *Result {
	ret := NewResult(coll, req, "error")
	ret.Error = err
	return ret
}
