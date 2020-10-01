package body

import "encoding/base64"

const KeyImage = "image"

type Image struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func NewImage(t string, bytes []byte) *Body {
	content := base64.StdEncoding.EncodeToString(bytes)
	return NewBody(KeyImage, &Image{Type: t, Content: content})
}

func (r *Image) ContentLength() int64 {
	return int64(len(r.Content))
}

func (r *Image) Bytes() []byte {
	b, _ := base64.StdEncoding.DecodeString(r.Content)
	return b
}

func (r *Image) MimeType() string {
	return r.Type
}

func (r *Image) String() string {
	return r.Content
}

func parseImage(contentType string, b []byte) *Body {
	return NewImage(contentType, b)
}
