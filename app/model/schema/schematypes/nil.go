package schematypes

const KeyNil = "nil"
type Nil struct {}

func (t Nil) Key() string {
	return KeyNil
}

func (t Nil) String() string {
	return t.Key()
}
