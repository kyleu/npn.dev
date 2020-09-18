package body

const KeyError = "error"

type Error struct {
	Message string `json:"message"`
}

func (l *Error) ContentLength() int64 {
	return int64(len(l.String()))
}

func (l *Error) Bytes() []byte {
	return []byte(l.String())
}

func (l *Error) MimeType() string {
	return ""
}

func (l *Error) String() string {
	return l.Message
}
