package schematypes

import (
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyTimestamp = "timestamp"

type Timestamp struct{}

func (t Timestamp) Key() string {
	return KeyTimestamp
}

func (t Timestamp) String() string {
	return t.Key()
}

func (t Timestamp) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
