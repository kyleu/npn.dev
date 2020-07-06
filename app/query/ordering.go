package query

type Ordering struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc,omitempty"`
}

func (o *Ordering) String() string {
	if o.Asc {
		return o.Column
	}
	return o.Column + "-desc"
}

type Orderings []*Ordering
