package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyChar = "char"

type Char struct{}

func (t Char) Key() string {
	return KeyChar
}

func (t Char) String() string {
	return t.Key()
}

func (t Char) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
