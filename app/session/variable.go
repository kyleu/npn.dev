package session

import (
	"fmt"
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
