package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyTime = "time"

type Time struct{}

func (t Time) Key() string {
	return KeyTime
}

func (t Time) String() string {
	return t.Key()
}

func (t Time) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
