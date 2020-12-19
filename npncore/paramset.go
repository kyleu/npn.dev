package npncore

import (
	"strings"

	"logur.dev/logur"
)

// A map of string keys to Params
type ParamSet map[string]*Params

// Gets the Params matching the provided key
func (s ParamSet) Get(key string, logger logur.Logger) *Params {
	x, ok := s[key]
	if !ok {
		return &Params{Key: key}
	}

	return x.Filtered(logger)
}

func (s ParamSet) String() string {
	ret := make([]string, 0, len(s))
	for _, p := range s {
		ret = append(ret, p.String())
	}

	return strings.Join(ret, ", ")
}
