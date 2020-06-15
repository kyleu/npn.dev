package schema

import "github.com/kyleu/npn/app/model/schema/schematypes"

type Field struct {
	Key      string              `json:"key"`
	Type     schematypes.Wrapped `json:"type"`
	Metadata *Metadata           `json:"metadata,omitempty"`
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
