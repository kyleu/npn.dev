package schema

type Paths []string

func (s Paths) Exists(key string) bool {
	for _, x := range s {
		if x == key {
			return true
		}
	}
	return false
}
