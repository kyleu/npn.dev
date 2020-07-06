package schematypes

import (
	"fmt"
	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/model/output"
)

const KeyList = "list"

type List struct {
	T Wrapped `json:"t"`
}

func (t List) Key() string {
	return KeyList
}

func (t List) String() string {
	return fmt.Sprintf("[]%v", t.T.String())
}

func (t List) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return fmt.Sprintf("[]%v", t.T.StringFor(ft, nr, src))
}
