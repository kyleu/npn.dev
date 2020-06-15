package schematypes

const KeyBool = "bool"
type Bool struct {}

func (t Bool) Key() string {
	return KeyBool
}

func (t Bool) String() string {
	return t.Key()
}

