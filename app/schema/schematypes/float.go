package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyFloat = "float"

type Float struct{}

func (t Float) Key() string {
	return KeyFloat
}

func (t Float) String() string {
	return t.Key()
}

func (t Float) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
