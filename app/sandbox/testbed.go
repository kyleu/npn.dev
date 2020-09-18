package sandbox

import (
	"github.com/kyleu/npn/npnweb"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		ret := "OK"
		return "Testbed!", ret, nil
	},
}
