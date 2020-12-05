package header

import (
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
	"net/http"
	"sort"
	"strings"
)

type Header struct {
	Key         string `json:"k,omitempty"`
	Value       string `json:"v,omitempty"`
	Description string `json:"desc,omitempty"`
}

func (h Header) Merge(data npncore.Data, logger logur.Logger) *Header {
	return &Header{
		Key:         npncore.MergeLog("header." + h.Key + ".key", h.Key, data, logger),
		Value:       npncore.MergeLog("header." + h.Key + ".value", h.Value, data, logger),
		Description: npncore.MergeLog("header." + h.Key + ".description", h.Description, data, logger),
	}
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

func (h Headers) Clone() Headers {
	ret := make(Headers, len(h))
	copy(ret, h)
	return ret
}

func (h Headers) Sort() {
	sort.SliceStable(h, func(l, r int) bool {
		return h[l].Key < h[r].Key
	})
}

func (h Headers) Set(k string, v string) Headers {
	matched := false
	hdr := &Header{Key: k, Value: v}
	for idx, x := range h {
		if strings.EqualFold(x.Key, k) {
			matched = true
			h[idx] = hdr
		}
	}
	if !matched {
		h = append(h, hdr)
	}
	return h
}

func (h Headers) Merge(data npncore.Data, logger logur.Logger) Headers {
	ret := make(Headers, 0, len(h))
	for _, x := range h {
		ret = append(ret, x.Merge(data, logger))
	}
	return ret
}
