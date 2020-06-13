package schema

type Method struct {
	Key      string    `json:"key"`
	Args     Fields    `json:"args,omitempty"`
	Ret      string    `json:"ret,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type Methods []*Method

func (s Methods) Get(key string) *Method {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
