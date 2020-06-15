package schematypes

const KeyReference = "reference"

type Reference struct {
	T string `json:"t"`
}

func (u Reference) Key() string {
	return KeyReference
}

func (u Reference) String() string {
	return u.T
}
