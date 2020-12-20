package session

import (
	"fmt"

	"github.com/kyleu/npn/npncore"
)

type Variable struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

func (v *Variable) String() string {
	return fmt.Sprintf("%v = %v", v.Key, v.Value)
}

func (v *Variable) Matches(x *Variable) bool {
	return v.Key == x.Key
}

func (v *Variable) Equals(x *Variable) bool {
	if v == nil || x == nil {
		return false
	}
	return *v == *x
}

type Variables []*Variable

func (v Variables) ToData() npncore.Data {
	ret := make(npncore.Data, len(v))
	for _, vr := range v {
		ret[vr.Key] = vr.Value
	}
	return ret
}
