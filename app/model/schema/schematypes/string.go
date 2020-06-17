package schematypes

import "fmt"

const KeyString = "string"
type String struct {
	MaxLength int `json:"maxLength,omitempty"`
}

func (s String) Key() string {
	return KeyString
}

func (s String) String() string {
	if s.MaxLength > 0 {
		return fmt.Sprintf("%v(%v)", s.Key(), s.MaxLength)
	}
	return s.Key()
}
