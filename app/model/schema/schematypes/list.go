package schematypes

const KeyList = "list"

type List struct {
	T Wrapped `json:"t"`
}

func (l List) Key() string {
	return KeyList
}

func (l List) String() string {
	return "[]" + l.T.String()
}
