package body

const KeyHTML = "html"

type HTML struct {
	Content string `json:"content"`
}

func NewHTML(content string) *Body {
	return NewBody(KeyHTML, &HTML{Content: content})
}

func (l *HTML) ContentLength() int64 {
	return int64(len(l.Content))
}

func (l *HTML) Bytes() []byte {
	return []byte(l.Content)
}

func (l *HTML) MimeType() string {
	return "text/html"
}

func (l *HTML) String() string {
	return l.Content
}

func parseHTML(b []byte) *Body {
	return NewHTML(string(b))
}
