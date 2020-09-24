package body

import "github.com/kyleu/npn/npncore"

const KeyJSON = "json"

type JSON struct {
	Msg interface{} `json:"msg"`
	str string
}

func NewJSON(msg interface{}) *Body {
	return &Body{Type: KeyJSON, Config: &JSON{Msg: msg}}
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
