package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyUnknown = "unknown"

type Unknown struct {
	X string `json:"x"`
}

func (t Unknown) Key() string {
	return KeyUnknown
}

func (t Unknown) String() string {
	return t.Key() + "(" + t.X + ")"
}

func (t Unknown) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
