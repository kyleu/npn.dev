package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyTimestampZoned = "timestampZoned"

type TimestampZoned struct{}

func (t TimestampZoned) Key() string {
	return KeyTimestampZoned
}

func (t TimestampZoned) String() string {
	return t.Key()
}

func (t TimestampZoned) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
