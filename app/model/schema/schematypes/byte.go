package schematypes

const KeyByte = "byte"
type Byte struct {}

func (t Byte) Key() string {
	return KeyByte
}

func (t Byte) String() string {
	return t.Key()
}

