package util

type Entry struct {
	K string
	V interface{}
}
type Entries []*Entry

func (e Entries) GetString(k string) string {
	for _, en := range e {
		if en.K == k {
			return en.V.(string)
		}
	}
	return ""
}
