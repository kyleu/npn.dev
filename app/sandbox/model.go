package sandbox

import "github.com/kyleu/npn/npnweb"

type Resolver func(ctx *npnweb.RequestContext) (string, interface{}, error)

type Sandbox struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	DevOnly     bool     `json:"devOnly,omitempty"`
	Resolve     Resolver `json:"-"`
}

type Sandboxes = []*Sandbox

var allSandboxes = Sandboxes{&Error}

func All() Sandboxes {
	return allSandboxes
}

func Register(s *Sandbox) *Sandbox {
	allSandboxes = append(allSandboxes, s)
	return s
}

func FromString(s string) *Sandbox {
	for _, t := range allSandboxes {
		if t.Key == s {
			return t
		}
	}
	return nil
}
