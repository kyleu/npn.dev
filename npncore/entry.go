package npncore

import (
	"fmt"
	"strings"
)

// a string key and interface{} value
type Entry struct {
	K string
	V interface{}
}

// a set of entries
type Entries []*Entry

// Returns a string representation of the value associated to the provided key
func (e Entries) GetString(k string) string {
	for _, en := range e {
		if en.K == k {
			return fmt.Sprint(en.V)
		}
	}
	return ""
}

// Returns a string array representation of the value associated to the provided key
func (e Entries) GetStringArray(k string) []string {
	return strings.Split(e.GetString(k), "||")
}

// Returns a new shallow copy of this Entries
func (e Entries) Clone() Entries {
	ret := make(Entries, 0, len(e))
	for _, en := range e {
		ret = append(ret, en)
	}
	return ret
}

// Returns a string representation of the contents of this Entries, mostly used for debugging
func (e Entries) String() string {
	ret := make([]string, 0, len(e))
	for _, p := range e {
		ret = append(ret, fmt.Sprintf("%v: %v", p.K, p.V))
	}
	return strings.Join(ret, ", ")
}
