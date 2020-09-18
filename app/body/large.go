package body

const KeyLarge = "large"

type Large struct {
	Filename string `json:"filename"`
	Length   int64  `json:"length"`
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
