package schematypes

const KeyUnknown = "unknown"

type Unknown struct {
	X string `json:"x"`
}

func (t Unknown) Key() string {
	return KeyUnknown
}

func (t Unknown) String() string {
	return t.Key() + ":" + t.X
}
