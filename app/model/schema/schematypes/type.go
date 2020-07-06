package schematypes

import (
	"fmt"
	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/model/output"
)

type Type interface {
	Key() string
	fmt.Stringer
	StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string
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

func (w Wrapped) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return w.V.StringFor(ft, nr, src)
}
