package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyOption = "option"

type Option struct {
	T Wrapped `json:"t"`
}

func (l Option) Key() string {
	return KeyOption
}

func (l Option) String() string {
	return "*" + l.T.String()
}

func (t Option) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return "*" + t.T.StringFor(ft, nr, src)
}
