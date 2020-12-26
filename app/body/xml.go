package body

import (
	"github.com/kyleu/libnpn/npncore"
	"logur.dev/logur"
)

const KeyXML = "xml"

type XML struct {
	Content string `json:"content"`
}

var _ Config = (*XML)(nil)

func NewXML(content string) *Body {
	return NewBody(KeyXML, &XML{Content: content})
}

func parseXML(b []byte) *Body {
	return NewXML(string(b))
}

func (h *XML) ContentLength() int64 {
	return int64(len(h.Content))
}

func (h *XML) Bytes() []byte {
	return []byte(h.Content)
}

func (h *XML) MimeType() string {
	return "text/xml"
}

func (h *XML) String() string {
	return h.Content
}

func (h *XML) Merge(data npncore.Data, logger logur.Logger) Config {
	return &XML{
		Content: npncore.MergeLog("body.xml.content", h.Content, data, logger),
	}
}

func (h *XML) Clone() *Body {
	return NewBody(KeyXML, &XML{Content: h.Content})
}
