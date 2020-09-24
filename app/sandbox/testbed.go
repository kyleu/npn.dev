package sandbox

import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npnweb"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		// return noop()
		return req(ctx)
	},
}

func noop() (string, interface{}, error) {
	ret := "Testbed!"
	return "Testbed", ret, nil
}

type reqrsp struct {
	Req   *request.Request
	Final header.Headers
	Rsp   *call.Result
}

func req(ctx *npnweb.RequestContext) (string, interface{}, error) {
	req, _ := request.FromString("sandbox", "http://localhost:10101/browse/debug/act/save")
	req.Prototype.Method = request.MethodPost
	// req.Prototype.Body = body.NewJSON(map[string]string{"a": "x", "b": "y", "c": "z"})
	req.Prototype.Body = body.NewForm(&body.FormEntry{K: "title", V: "DEBUG!"})
	// req.Prototype.Auth = auth.Auths{auth.NewBasic("kyle", "kyleu", true)}

	rsp := app.Svc(ctx.App).Caller.Call(req.Prototype)

	ret := reqrsp{Req: req, Final: req.Prototype.FinalHeaders(), Rsp: rsp}

	return "Request", ret, nil
}

/*
func callJS(ctx *npnweb.RequestContext) (string, interface{}, error) {
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
		return "", ret, err
	}
	return "JavaScript", ret, nil
}

func callLua(ctx *npnweb.RequestContext) (string, interface{}, error) {
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
		return "", ret, err
	}
	return "Lua", ret, nil
}
*/
