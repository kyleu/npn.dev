package schematypes

const KeyDate = "date"
type Date struct {}

func (t Date) Key() string {
	return KeyDate
}

func (t Date) String() string {
	return t.Key()
}
