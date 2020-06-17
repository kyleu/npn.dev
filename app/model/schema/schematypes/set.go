package schematypes

import "fmt"

const KeySet = "set"
type Set struct {
	T Wrapped
}

func (s Set) Key() string {
	return KeySet
}

func (s Set) String() string {
	return fmt.Sprintf("%v[%v]", s.Key(), s.T.String())
}
