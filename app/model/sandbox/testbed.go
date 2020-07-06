package sandbox

import (
	"github.com/kyleu/npn/app/model/loader/cloudbeaver"
	"github.com/kyleu/npn/app/web"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx *web.RequestContext) (string, interface{}, error) {
		loader := cloudbeaver.NewLoader(false, "localhost", 8978, ctx.Logger)
		detect, err := loader.Detect()
		if err != nil {
			return "", nil, err
		}
		sch, data, err := loader.Crawl()
		if err != nil {
			return "", nil, err
		}
		ret := map[string]interface{}{
			"detect": detect,
			"schema": sch,
			"data":   data,
		}
		return "Testbed!", ret, nil
	},
}
