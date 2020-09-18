package body

import "github.com/kyleu/npn/npncore"

const KeyJSON = "json"

type JSON struct {
	Msg    interface{} `json:"msg"`
	Length int64       `json:"length"`
}

func (j *JSON) ContentLength() int64 {
	return j.Length
}

func (j *JSON) Bytes() []byte {
	return []byte(j.String())
}

func (j *JSON) MimeType() string {
	return "application/json"
}

func (j *JSON) String() string {
	return string(npncore.ToJSON(j.Msg, nil))
}
