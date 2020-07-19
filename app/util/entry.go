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

func (e Entries) Clone() Entries {
	ret := make(Entries, 0, len(e))
	for _, en := range e {
		ret = append(ret, en)
	}
	return ret
}
