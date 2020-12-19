package body

import (
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

const KeyHTML = "html"

type HTML struct {
	Content string `json:"content"`
}

var _ Config = (*HTML)(nil)

func NewHTML(content string) *Body {
	return NewBody(KeyHTML, &HTML{Content: content})
}

func parseHTML(b []byte) *Body {
	return NewHTML(string(b))
}

func (h *HTML) ContentLength() int64 {
	return int64(len(h.Content))
}

func (h *HTML) Bytes() []byte {
	return []byte(h.Content)
}

func (h *HTML) MimeType() string {
	return "text/html"
}

func (h *HTML) String() string {
	return h.Content
}

func (h *HTML) Merge(data npncore.Data, logger logur.Logger) Config {
	return &HTML{
		Content: npncore.MergeLog("body.html.content", h.Content, data, logger),
	}
}

func (h *HTML) Clone() *Body {
	return NewBody(KeyHTML, &HTML{Content: h.Content})
}
