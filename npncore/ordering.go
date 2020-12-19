package npncore

// Represents a column name and a boolean indicating sort order
type Ordering struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

// Query string format
func (o Ordering) String() string {
	if o.Asc {
		return o.Column
	}
	return o.Column + "-desc"
}

// Helper for Ordering arrays
type Orderings []*Ordering

// Default ordering, by "created" descending
var DefaultCreatedOrdering = Orderings{{Column: "created", Asc: false}}

// Empty Orderings
var NoOrdering = Orderings{}
