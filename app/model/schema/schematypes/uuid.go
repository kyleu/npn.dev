package schematypes

const KeyUUID = "uuid"
type UUID struct {}

func (t UUID) Key() string {
	return KeyUUID
}

func (t UUID) String() string {
	return t.Key()
}
