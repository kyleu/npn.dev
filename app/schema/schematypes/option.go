package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyOption = "option"

type Option struct {
	T Wrapped `json:"t"`
}

func (o Option) Key() string {
	return KeyOption
}

func (o Option) String() string {
	return "*" + o.T.String()
}

func (o Option) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return "*" + o.T.StringFor(ft, nr, src)
}
