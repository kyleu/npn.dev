package body

const KeyTemp = "temp"

type Temp struct {
	Foo string `json:"foo"`
}

func (t *Temp) Bytes() []byte {
	return []byte("TODO")
}

func (t *Temp) String() string {
	return t.Foo
}
