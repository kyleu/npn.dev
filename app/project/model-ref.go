package project

import (
	"reflect"

	"github.com/kyleu/npn/app/util"
)

type ModelRef struct {
	Pkg         util.Pkg `json:"pkg"`
	Key         string   `json:"key"`
	PkgOverride util.Pkg `json:"pkgOverride,omitempty"`
}

type ModelRefs []*ModelRef

func (m ModelRefs) Get(pkg util.Pkg, key string) *ModelRef {
	for _, x := range m {
		if reflect.DeepEqual(x.Pkg, pkg) && x.Key == key {
			return x
		}
	}
	return nil
}
