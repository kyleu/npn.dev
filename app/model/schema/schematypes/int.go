package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyInt = "int"

type Int struct{}

func (t Int) Key() string {
	return KeyInt
}

func (t Int) String() string {
	return t.Key()
}

func (t Int) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
