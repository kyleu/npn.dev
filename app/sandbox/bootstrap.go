package sandbox

import (
	"github.com/kyleu/npn/npnweb"
)

var Bootstrap = Register(&Sandbox{
	Key:         "bootstrap",
	Title:       "Bootstrap",
	Description: "Packages the bootstrap projects for release",
	DevOnly:     true,
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		return "OK", "TODO", nil
	},
})
