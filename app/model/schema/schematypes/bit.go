package schematypes

const KeyBit = "bit"
type Bit struct {}

func (t Bit) Key() string {
	return KeyBit
}

func (t Bit) String() string {
	return t.Key()
}

