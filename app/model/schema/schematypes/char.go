package schematypes

const KeyChar = "char"
type Char struct {}

func (t Char) Key() string {
	return KeyChar
}

func (t Char) String() string {
	return t.Key()
}
