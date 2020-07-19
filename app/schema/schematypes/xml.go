package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyXML = "xml"

type XML struct{}

func (t XML) Key() string {
	return KeyXML
}

func (t XML) String() string {
	return t.Key()
}

func (t XML) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
