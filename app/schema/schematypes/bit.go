package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyBit = "bit"

type Bit struct{}

func (t Bit) Key() string {
	return KeyBit
}

func (t Bit) String() string {
	return t.Key()
}

func (t Bit) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
