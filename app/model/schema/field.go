package schema

type Field struct {
	Key string `json:"key"`
	Type string `json:"type"`
	Optional bool `json:"optional,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type Fields []*Field

func (s Fields) Get(key string) *Field {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
