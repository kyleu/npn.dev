package schematypes

import (
	"fmt"
	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/model/output"
)

const KeyString = "string"

type String struct {
	MaxLength int    `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
}

func (s String) Key() string {
	return KeyString
}

func (s String) String() string {
	if s.MaxLength > 0 {
		return fmt.Sprintf("%v(%v)", s.Key(), s.MaxLength)
	}
	return s.Key()
}

func (t String) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	return nr.Get(nil, t.Key(), src).String()
}
