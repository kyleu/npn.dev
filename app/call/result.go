package call

import (
	"github.com/gofrs/uuid"
)

type Result struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Collection string    `json:"collection,omitempty"`
	Request    string    `json:"request,omitempty"`
	Status     string    `json:"status,omitempty"`
	Response   *Response `json:"response,omitempty"`
	Error      string    `json:"error,omitempty"`
}

func NewResult(id uuid.UUID, coll string, req string, status string) *Result {
	return &Result{ID: id, Collection: coll, Request: req, Status: status}
}

func NewErrorResult(id uuid.UUID, coll string, req string, err string) *Result {
	ret := NewResult(id, coll, req, "error")
	ret.Error = err
	return ret
}
