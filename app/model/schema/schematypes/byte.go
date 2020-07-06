package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyByte = "byte"

type Byte struct{}

func (t Byte) Key() string {
	return KeyByte
}

func (t Byte) String() string {
	return t.Key()
}

func (t Byte) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
