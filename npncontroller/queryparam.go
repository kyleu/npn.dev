package npncontroller

import (
	"net/url"
	"strings"

	"github.com/kyleu/npn/npncore"
)

type QueryParam struct {
	Key         string `json:"k,omitempty"`
	Value       string `json:"v,omitempty"`
	Description string `json:"desc,omitempty"`
}

func (q *QueryParam) String() string {
	return url.QueryEscape(q.Key) + "=" + url.QueryEscape(q.Value)
}

type QueryParams []*QueryParam

func (q QueryParams) String() string {
	ret := make([]string, 0, len(q))
	for _, x := range q {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, "&")
}

func QueryParamsFromRaw(s string) QueryParams {
	ret := make(QueryParams, 0)
	parts := strings.Split(s, "&")
	for _, x := range parts {
		if len(x) > 0 {
			k, v := npncore.SplitString(x, '=', true)
			k, _ = url.QueryUnescape(k)
			v, _ = url.QueryUnescape(v)
			ret = append(ret, &QueryParam{Key: k, Value: v})
		}
	}
	return ret
}
