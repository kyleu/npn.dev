package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyEnumValue = "enumValue"

type EnumValue struct{}

func (t EnumValue) Key() string {
	return KeyEnumValue
}

func (t EnumValue) String() string {
	return t.Key()
}

func (t EnumValue) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
