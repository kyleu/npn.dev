package schematypes

import "fmt"

const KeyMap = "map"
type Map struct {
	K Wrapped
	V Wrapped
}

func (t Map) Key() string {
	return KeyMap
}

func (t Map) String() string {
	return fmt.Sprintf("%v[%v]%v", t.Key(), t.K.String(), t.V.String())
}
