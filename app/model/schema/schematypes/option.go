package schematypes

const KeyOption = "option"

type Option struct {
	T Wrapped `json:"t"`
}

func (l Option) Key() string {
	return KeyOption
}

func (l Option) String() string {
	return "*" + l.T.String()
}
