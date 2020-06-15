package schematypes

const KeyString = "string"
type String struct {}

func (t String) Key() string {
	return KeyString
}

func (t String) String() string {
	return t.Key()
}
