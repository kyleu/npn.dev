package body

const KeyHTML = "HTML"

type HTML struct {
	Content string `json:"content"`
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
