package session

import (
	"fmt"
)

type Variable struct {
	Key     string    `json:"k"`
	Value    string    `json:"v"`
}

func (v *Variable) String() string {
	return fmt.Sprintf("%v = %v", v.Key, v.Value)
}

type Variables []*Variable
