package schema

type Union struct {
	Key      string   `json:"key"`
	Variants []*Field `json:"variants"`

	Metadata *Metadata `json:"metadata,omitempty"`
}

type Unions []*Union

func (s Unions) Get(key string) *Union {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
