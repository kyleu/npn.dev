package body

import (
	"encoding/base64"
	"github.com/sirupsen/logrus"

	"github.com/kyleu/libnpn/npncore"
)

const KeyImage = "image"

type Image struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

var _ Config = (*Image)(nil)

func NewImage(t string, bytes []byte) *Body {
	content := base64.StdEncoding.EncodeToString(bytes)
	return NewBody(KeyImage, &Image{Type: t, Content: content})
}

func parseImage(contentType string, b []byte) *Body {
	return NewImage(contentType, b)
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

func (r *Image) Merge(data npncore.Data, logger *logrus.Logger) Config {
	return &Image{
		Type:    npncore.MergeLog("body.image.type", r.Type, data, logger),
		Content: npncore.MergeLog("body.image.content", r.Content, data, logger),
	}
}

func (r *Image) Clone() *Body {
	return NewBody(KeyImage, &Image{Type: r.Type, Content: r.Content})
}
