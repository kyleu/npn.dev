package header

import "strings"

type Header struct {
	Key   string `json:"k,omitempty"`
	Value string `json:"v,omitempty"`
}

type Headers []*Header

func (h Headers) Contains(k string) bool {
	for _, x := range h {
		if strings.EqualFold(x.Key, k) {
			return true
		}
	}
	return false
}

func (h Headers) Get(k string) *Header {
	k = strings.ToLower(k)
	for _, x := range h {
		if strings.EqualFold(x.Key, k) {
			return x
		}
	}
	return nil
}

func (h Headers) GetValue(k string) string {
	x := h.Get(k)
	if x == nil {
		return ""
	}
	return x.Value
}
