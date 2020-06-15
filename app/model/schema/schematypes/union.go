package schematypes

import "strings"

const KeyUnion = "union"

type Union struct {
	Variants []Wrapped `json:"variants"`
}

func (u Union) Key() string {
	return KeyUnion
}

func (u Union) String() string {
	ss := make([]string, 0, len(u.Variants))
	for _, variant := range u.Variants {
		ss = append(ss, variant.String())
	}
	return strings.Join(ss, " | ")
}
