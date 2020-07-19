package schematypes

import (
	"fmt"
	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/output"
)

const KeySet = "set"

type Set struct {
	T Wrapped
}

func (t Set) Key() string {
	return KeySet
}

func (t Set) String() string {
	return fmt.Sprintf("%v[%v]", t.Key(), t.T.String())
}

func (t Set) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return fmt.Sprintf("%v[%v]", t.Key(), t.T.StringFor(ft, nr, src))
}
