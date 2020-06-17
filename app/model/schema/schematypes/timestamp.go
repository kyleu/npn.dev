package schematypes

const KeyTimestamp = "timestamp"
type Timestamp struct {}

func (t Timestamp) Key() string {
	return KeyTimestamp
}

func (t Timestamp) String() string {
	return t.Key()
}
