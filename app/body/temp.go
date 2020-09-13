package body

const KeyTemp = "temp"

type Temp struct {
	Foo string `json:"foo"`
}

func (t *Temp) ContentLength() int64 {
	return int64(len([]byte(t.Foo)))
}

func (t *Temp) Bytes() []byte {
	return []byte(t.Foo)
}

func (t *Temp) MimeType() string {
	return "text/plain"
}

func (t *Temp) String() string {
	return t.Foo
}
