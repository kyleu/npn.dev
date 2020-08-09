package npncore

import (
	"fmt"
	"strings"
)

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

func (e Entries) GetStringArray(k string) []string {
	for _, en := range e {
		if en.K == k {
			return strings.Split(en.V.(string), "||")
		}
	}
	return nil
}

func (e Entries) Clone() Entries {
	ret := make(Entries, 0, len(e))
	for _, en := range e {
		ret = append(ret, en)
	}
	return ret
}

func (e Entries) String() string {
	ret := make([]string, 0, len(e))
	for _, p := range e {
		ret = append(ret, fmt.Sprintf("%v: %v", p.K, p.V))
	}
	return strings.Join(ret, ", ")
}
