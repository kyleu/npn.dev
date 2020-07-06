package schematypes

import (
	"fmt"
	"github.com/kyleu/npn/app/util"
	"strings"

	"github.com/kyleu/npn/app/model/output"
)

type Argument struct {
	Key  string  `json:"key"`
	Type Wrapped `json:"type"`
}

func (a Argument) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	t := nr.Get(nil, a.Type.Key(), src)
	return fmt.Sprintf("%v %v", a.Key, t.String())
}

func (a Argument) String() string {
	return fmt.Sprintf("%v %v", a.Key, a.Type.String())
}

type Arguments []Argument

const KeyMethod = "method"

type Method struct {
	Args Arguments `json:"args,omitempty"`
	Ret  Wrapped   `json:"ret,omitempty"`
}

func (t Method) Key() string {
	return KeyMethod
}

func (t Method) String() string {
	argStrings := make([]string, 0, len(t.Args))
	for _, arg := range t.Args {
		argStrings = append(argStrings, arg.String())
	}
	return fmt.Sprintf("fn(%v) %v", strings.Join(argStrings, ", "), t.Ret.String())
}

func (t Method) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	argStrings := make([]string, 0, len(t.Args))
	for _, arg := range t.Args {
		argStrings = append(argStrings, arg.StringFor(ft, nr, src))
	}
	return fmt.Sprintf("fn(%v) %v", strings.Join(argStrings, ", "), t.Ret.StringFor(ft, nr, src))
}
