package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyJSON = "json"

type JSON struct{}

func (t JSON) Key() string {
	return KeyJSON
}

func (t JSON) String() string {
	return t.Key()
}

func (t JSON) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
