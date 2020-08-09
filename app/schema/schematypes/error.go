package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyError = "error"

type Error struct {
	Message string `json:"t"`
}

func (e Error) Key() string {
	return KeyError
}

func (e Error) String() string {
	return "error(" + e.Message + ")"
}

func (e Error) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, e.Key(), src).String()
}
