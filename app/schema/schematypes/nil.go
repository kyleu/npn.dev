package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyNil = "nil"

type Nil struct{}

func (t Nil) Key() string {
	return KeyNil
}

func (t Nil) String() string {
	return t.Key()
}

func (t Nil) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
