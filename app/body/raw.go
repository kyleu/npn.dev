package body

import (
	"encoding/base64"
	"unicode/utf8"
)

const KeyRaw = "raw"

type Raw struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Length  int64 `json:"length,omitempty"`
	Binary  bool `json:"binary,omitempty"`
}

func NewRaw(bytes []byte) *Body {
	cfg := &Raw{Length: int64(len(bytes)), Binary: true}
	if(utf8.Valid(bytes)) {
		cfg.Content = string(bytes)
	} else {
		cfg.Content = base64.StdEncoding.EncodeToString(bytes)
		cfg.Binary = true
	}
	return &Body{Type: KeyRaw, Config: cfg}
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
