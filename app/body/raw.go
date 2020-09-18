package body

const KeyRaw = "raw"

type Raw struct {
	bytes []byte `json:"bytes"`
}

func (r *Raw) ContentLength() int64 {
	return int64(len(r.bytes))
}

func (r *Raw) Bytes() []byte {
	return r.bytes
}

func (r *Raw) MimeType() string {
	return "???"
}

func (r *Raw) String() string {
	return string(r.bytes)
}
