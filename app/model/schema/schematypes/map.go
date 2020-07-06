package schematypes

import (
	"fmt"
	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/model/output"
)

const KeyMap = "map"

type Map struct {
	K Wrapped
	V Wrapped
}

func (t Map) Key() string {
	return KeyMap
}

func (t Map) String() string {
	return fmt.Sprintf("%v[%v]%v", t.Key(), t.K.String(), t.V.String())
}

func (t Map) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	k := t.K.StringFor(ft, nr, src)
	v := t.V.StringFor(ft, nr, src)
	return fmt.Sprintf("%v[%v]%v", t.Key(), k, v)
}
