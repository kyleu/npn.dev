package header

type Header struct {
	Key   string `json:"k,omitempty"`
	Value string `json:"v,omitempty"`
}

type Headers []*Header
