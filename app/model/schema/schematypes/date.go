package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyDate = "date"

type Date struct{}

func (t Date) Key() string {
	return KeyDate
}

func (t Date) String() string {
	return t.Key()
}

func (t Date) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
