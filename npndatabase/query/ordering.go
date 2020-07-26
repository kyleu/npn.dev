package query

import (
	"github.com/fevo-tech/charybdis/app/util"
)

type Ordering struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

func (o Ordering) String() string {
	if o.Asc {
		return o.Column
	}
	return o.Column + "-desc"
}

type Orderings []*Ordering

var DefaultCreatedOrdering = Orderings{{Column: util.KeyCreated, Asc: false}}
var NoOrdering = Orderings{}
