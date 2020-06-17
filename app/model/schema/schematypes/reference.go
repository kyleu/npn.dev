package schematypes

import "strings"

const KeyReference = "reference"

type Reference struct {
	Pkg []string `json:"pkg,omitempty"`
	T   string `json:"t"`
}

func (u Reference) Key() string {
	return KeyReference
}

func (u Reference) String() string {
	if len(u.Pkg) == 0 {
		return u.T
	}
	return strings.Join(append(u.Pkg, u.T), ".")
}
