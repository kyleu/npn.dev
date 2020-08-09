package sandbox

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
)

var Error = Sandbox{
	Key:         npncore.KeyError,
	Title:       npncore.Title(npncore.KeyError),
	Description: "An example of the error page",
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		return "", nil, errors.New("here's an intentional error")
	},
}
