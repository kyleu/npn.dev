package schematypes

import (
	"fmt"
)

type Type interface {
	Key() string
	fmt.Stringer
}

type Wrapped struct {
	K string `json:"k"`
	V Type   `json:"t,omitempty"`
}

func Wrap(t Type) Wrapped {
	_, ok := t.(Wrapped)
	if ok {
		return t.(Wrapped)
	}
	return Wrapped{K: t.Key(), V: t}
}

func (w Wrapped) Key() string {
	return w.K
}

func (w Wrapped) String() string {
	return w.V.String()
}
