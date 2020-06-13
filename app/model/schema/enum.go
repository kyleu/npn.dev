package schema

type EnumValue struct {
	Key      string    `json:"key"`
	IntVal   int       `json:"intVal,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type EnumValues []*EnumValue

func (s EnumValues) Get(key string) *EnumValue {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (s EnumValues) GetInt(key int) *EnumValue {
	for _, x := range s {
		if x.IntVal == key {
			return x
		}
	}
	return nil
}

type Enum struct {
	Key      string       `json:"key"`
	Values   []*EnumValue `json:"values"`
	Metadata *Metadata    `json:"metadata"`
}

type Enums []*Enum

func (s Enums) Get(key string) *Enum {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
