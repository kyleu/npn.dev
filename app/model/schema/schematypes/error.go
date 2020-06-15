package schematypes

const KeyError = "error"

type Error struct {
	Message string `json:"t"`
}

func (e Error) Key() string {
	return KeyError
}

func (e Error) String() string {
	return "error(`" + e.Message + "`)"
}
