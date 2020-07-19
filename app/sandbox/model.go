package sandbox

import "github.com/kyleu/npn/app/web"

type Resolver func(ctx *web.RequestContext) (string, interface{}, error)

type Sandbox struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	DevOnly     bool     `json:"devOnly,omitempty"`
	Resolve     Resolver `json:"-"`
}

type Sandboxes = []*Sandbox

var AllSandboxes = Sandboxes{&Bootstrap, &Testbed, &Error}

func FromString(s string) *Sandbox {
	for _, t := range AllSandboxes {
		if t.Key == s {
			return t
		}
	}
	return nil
}
