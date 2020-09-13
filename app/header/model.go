package header

import (
	"net/http"
	"strings"
)

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

func (h Headers) ToHTTP() http.Header {
	ret := make(http.Header, len(h))
	for _, hd := range h {
		ret[hd.Key] = append(ret[hd.Key], hd.Value)
	}
	return ret
}
