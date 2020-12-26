package body

import (
	"github.com/kyleu/libnpn/npncore"
	"logur.dev/logur"
)

const KeyJSON = "json"

type JSON struct {
	Msg interface{} `json:"msg"`
	str string
}

var _ Config = (*JSON)(nil)

func NewJSON(msg interface{}) *Body {
	return NewBody(KeyJSON, &JSON{Msg: msg})
}

func parseJSON(ct string, charset string, b []byte) *Body {
	var x interface{}
	err := npncore.FromJSON(b, &x)
	if err != nil {
		if ct == "" {
			return NewError(err.Error())
		}
		return detect("", "", charset, b)
	}

	return NewJSON(x)
}

func (j *JSON) ContentLength() int64 {
	return int64(len(j.Bytes()))
}

func (j *JSON) Bytes() []byte {
	return []byte(j.String())
}

func (j *JSON) MimeType() string {
	return "application/json"
}

func (j *JSON) String() string {
	if len(j.str) == 0 {
		j.str = npncore.ToJSON(j.Msg, nil)
	}
	return j.str
}

func (j *JSON) Merge(data npncore.Data, logger logur.Logger) Config {
	m := j.Msg
	ms := npncore.ToJSONCompact(j.Msg, nil)
	if len(ms) > 0 {
		ms = npncore.MergeLog("body.json.msg", ms, data, logger)
		var i interface{}
		err := npncore.FromJSON([]byte(ms), &i)
		if err == nil && i != nil {
			m = i
		}
	}

	s := j.str
	if len(s) > 0 {
		s = npncore.MergeLog("body.json.str", s, data, logger)
	}
	return &JSON{Msg: m, str: s}
}

func (j *JSON) Clone() *Body {
	return NewBody(KeyJSON, &JSON{Msg: j.Msg, str: j.str})
}
