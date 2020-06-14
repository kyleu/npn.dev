package schema

type Scalar struct {
	Key      string    `json:"key"`
	Type     string    `json:"type"`
	Metadata *Metadata `json:"metadata"`
}

type Scalars []*Scalar

func (s Scalars) Get(key string) *Scalar {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
