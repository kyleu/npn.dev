package sandbox

import (
	"github.com/kyleu/npn/npnweb"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve:     noop,
}

func noop(ctx *npnweb.RequestContext) (string, interface{}, error) {
	ret := "Testbed!"
	return "Testbed", ret, nil
}

/* Req
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
*/

/* JavaScript
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
*/

/* Lua
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
