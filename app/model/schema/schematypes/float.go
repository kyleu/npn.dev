package schematypes

const KeyFloat = "float"
type Float struct {}

func (t Float) Key() string {
	return KeyFloat
}

func (t Float) String() string {
	return t.Key()
}
