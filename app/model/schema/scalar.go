package schema

import "github.com/kyleu/npn/app/util"

type Scalar struct {
	Pkg      []string    `json:"pkg"`
	Key      string    `json:"key"`
	Type     string    `json:"type"`
	Metadata *Metadata `json:"metadata"`
}

type Scalars []*Scalar

func (s Scalars) Get(pkg []string, key string) *Scalar {
	for _, x := range s {
		if util.StringArraysEqual(x.Pkg, pkg) && x.Key == key {
			return x
		}
	}
	return nil
}
