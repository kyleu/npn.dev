package schema

import "emperror.dev/errors"

type Model struct {
	Key      string    `json:"key"`
	Fields   Fields    `json:"fields"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

func (m *Model) AddField(f *Field) error {
	if m.Fields.Get(f.Key) != nil {
		return errors.New("field [" + f.Key + "] already exists")
	}
	m.Fields = append(m.Fields, f)
	return nil
}

type Models []*Model

func (s Models) Get(key string) *Model {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
