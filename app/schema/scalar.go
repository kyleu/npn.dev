package schema

import (
	"reflect"

	"github.com/kyleu/npn/app/util"
)

type Scalar struct {
	Pkg         util.Pkg  `json:"pkg,omitempty"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Description string    `json:"description,omitempty"`
	Metadata    *Metadata `json:"metadata,omitempty"`
}

type Scalars []*Scalar

func (s Scalars) Get(pkg util.Pkg, key string) *Scalar {
	for _, x := range s {
		if reflect.DeepEqual(x.Pkg, pkg) && x.Key == key {
			return x
		}
	}
	return nil
}
