package schematypes

const KeyInt = "int"
type Int struct {}

func (t Int) Key() string {
	return KeyInt
}

func (t Int) String() string {
	return t.Key()
}
