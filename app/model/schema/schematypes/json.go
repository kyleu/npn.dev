package schematypes

const KeyJSON = "json"
type JSON struct {}

func (t JSON) Key() string {
	return KeyJSON
}

func (t JSON) String() string {
	return t.Key()
}
