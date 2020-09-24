package sandbox

import (
	"github.com/kyleu/npn/npnweb"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		ret, err := noop()
		return "Testbed!", ret, err
	},
}

func noop() (string, error) {
	return "Testbed!", nil
}

/*
func callJS(ctx *npnweb.RequestContext) (interface{}, error) {
	jssvc := js.NewService(ctx.Logger)

	req := request.Request{
		Key:         "xkey",
		Title:       "xtitle",
		Description: "xdesc",
		Prototype:   request.PrototypeFromString("https://google.com/q?x=1"),
	}
	jssvc.Set("req", req)

	ret, err := jssvc.Call("req.prototype.domain")
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func callLua(ctx *npnweb.RequestContext) (interface{}, error) {
	luasvc := lua.NewService(ctx.Logger)

	req := request.Request{
		Key:         "xkey",
		Title:       "xtitle",
		Description: "xdesc",
		Prototype:   request.PrototypeFromString("https://google.com/q?x=1"),
	}
	luasvc.Set("req", req)

	ret, err := luasvc.Call("print(111)")
	if err != nil {
		return ret, err
	}
	return ret, nil
}
*/
