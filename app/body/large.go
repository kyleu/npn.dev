package body

const KeyLarge = "large"

type Large struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Length      int64  `json:"length"`
}

var _ Config = (*Large)(nil)

func NewLarge(filename string, contentType string, length int64) *Body {
	return NewBody(KeyLarge, &Large{Filename: filename, ContentType: contentType, Length: length})
}

func (l *Large) ContentLength() int64 {
	return l.Length
}

func (l *Large) Bytes() []byte {
	return []byte(l.String())
}

func (l *Large) MimeType() string {
	return ""
}

func (l *Large) String() string {
	return l.Filename
}
