package schematypes

import (
	"strings"

	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/util"
)

const KeyReference = "reference"

type Reference struct {
	Pkg util.Pkg `json:"pkg,omitempty"`
	T   string   `json:"t"`
}

func (t Reference) Key() string {
	return KeyReference
}

func (t Reference) String() string {
	return strings.Join(append(t.Pkg, t.T), ".")
}

func (t Reference) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(t.Pkg, t.T, src).Class()
}
