package sandbox

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/app/web"
)

var Error = Sandbox{
	Key:         util.KeyError,
	Title:       util.Title(util.KeyError),
	Description: "An example of the error page",
	Resolve: func(ctx *web.RequestContext) (string, interface{}, error) {
		return "", nil, errors.New("here's an intentional error")
	},
}
