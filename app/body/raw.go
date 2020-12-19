package body

import (
	"encoding/base64"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
	"unicode/utf8"
)

const KeyRaw = "raw"

type Raw struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Length  int64  `json:"length,omitempty"`
	Binary  bool   `json:"binary,omitempty"`
}

var _ Config = (*Raw)(nil)

func NewRaw(bytes []byte) *Body {
	cfg := &Raw{Length: int64(len(bytes)), Binary: true}
	if utf8.Valid(bytes) {
		cfg.Content = string(bytes)
	} else {
		cfg.Content = base64.StdEncoding.EncodeToString(bytes)
		cfg.Binary = true
	}
	return NewBody(KeyRaw, cfg)
}

func (r *Raw) ContentLength() int64 {
	return r.Length
}

func (r *Raw) Bytes() []byte {
	if r.Binary {
		b, _ := base64.StdEncoding.DecodeString(r.Content)
		return b
	}
	return []byte(r.Content)
}

func (r *Raw) MimeType() string {
	return r.Type
}

func (r *Raw) String() string {
	return r.Content
}

func (r *Raw) Merge(data npncore.Data, logger logur.Logger) Config {
	return &Raw{
		Type:    npncore.MergeLog("body.raw.type", r.Type, data, logger),
		Content: npncore.MergeLog("body.raw.content", r.Content, data, logger),
		Length:  r.Length,
		Binary:  r.Binary,
	}
}

func (r *Raw) Clone() *Body {
	return NewBody(KeyRaw, &Raw{Type: r.Type, Content: r.Content, Length: r.Length, Binary: r.Binary})
}
