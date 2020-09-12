package request

type Body interface {
	Type() string
	Bytes() []byte
}
