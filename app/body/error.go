package body

import (
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

const KeyError = "error"

type Error struct {
	Message string `json:"message"`
}

var _ Config = (*Error)(nil)

func NewError(message string) *Body {
	return NewBody(KeyError, &Error{Message: message})
}

func (e *Error) ContentLength() int64 {
	return int64(len(e.String()))
}

func (e *Error) Bytes() []byte {
	return []byte(e.String())
}

func (e *Error) MimeType() string {
	return "text/plain"
}

func (e *Error) String() string {
	return e.Message
}

func (e *Error) Merge(data npncore.Data, logger logur.Logger) Config {
	return &Error{Message: npncore.MergeLog("body.error.message", e.Message, data, logger)}
}

