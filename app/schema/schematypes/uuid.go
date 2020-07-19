package schematypes

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/util"
)

const KeyUUID = "uuid"

type UUID struct{}

func (t UUID) Key() string {
	return KeyUUID
}

func (t UUID) String() string {
	return t.Key()
}

func (t UUID) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
