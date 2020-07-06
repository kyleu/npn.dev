package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyBool = "bool"

type Bool struct{}

func (t Bool) Key() string {
	return KeyBool
}

func (t Bool) String() string {
	return t.Key()
}

func (t Bool) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
