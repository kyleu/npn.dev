package schematypes

import (
	"strings"
)

type Argument struct {
	Key  string  `json:"key"`
	Type Wrapped `json:"type"`
}

func (a Argument) String() string {
	return a.Key + ": " + a.Type.String()
}

type Arguments []Argument

const KeyMethod = "method"

type Method struct {
	Args Arguments
	Ret  Wrapped
}

func (t Method) Key() string {
	return KeyMethod
}

func (t Method) String() string {
	argStrings := make([]string, 0, len(t.Args))
	for _, arg := range t.Args {
		argStrings = append(argStrings, arg.Key + ": " + arg.Type.String())
	}
	return "(" + strings.Join(argStrings, ", ") + ")" + ": " + t.Ret.String()
}
