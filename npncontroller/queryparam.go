package npncontroller

import (
	"logur.dev/logur"
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

func (q *QueryParam) Merge(data npncore.Data, logger logur.Logger) *QueryParam {
	return &QueryParam{
		Key:         npncore.MergeLog("query." + q.Key + ".key", q.Key, data, logger),
		Value:       npncore.MergeLog("query." + q.Key + ".value", q.Value, data, logger),
		Description: npncore.MergeLog("query." + q.Key + ".description", q.Description, data, logger),
	}
}

type QueryParams []*QueryParam

func (q QueryParams) String() string {
	ret := make([]string, 0, len(q))
	for _, x := range q {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, "&")
}

func (q QueryParams) Merge(data npncore.Data, logger logur.Logger) QueryParams {
	ret := make(QueryParams, 0, len(q))
	for _, qp := range q {
		ret = append(ret, qp.Merge(data, logger))
	}
  return ret
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
