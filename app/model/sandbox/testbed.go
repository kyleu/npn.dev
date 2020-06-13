package sandbox

import (
	"github.com/kyleu/npn/app/web"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx *web.RequestContext) (string, interface{}, error) {
		var rsp = ""
		return "Testbed!", rsp, nil
	},
}
