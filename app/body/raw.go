package body

const KeyRaw = "raw"

type Raw struct {
	Content []byte `json:"content"`
}

func NewRaw(bytes []byte) *Body {
	return &Body{Type: KeyRaw, Config: &Raw{Content: bytes}}
}

func (r *Raw) ContentLength() int64 {
	return int64(len(r.Content))
}

func (r *Raw) Bytes() []byte {
	return r.Content
}

func (r *Raw) MimeType() string {
	return "application/octet-stream"
}

func (r *Raw) String() string {
	return string(r.Content)
}
