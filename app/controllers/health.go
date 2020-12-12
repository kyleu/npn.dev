package controllers

import (
	"github.com/kyleu/npn/npncore"
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
)

func Health(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		x := map[string]interface{}{
			"status": "OK",
		}
		_, _ = w.Write(npncore.ToJSONBytes(x, ctx.Logger, true))
		return "", nil
	})
}

func Testbed(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ret := "Hello from testbed!"
		_, _ = w.Write([]byte(ret))
		return "", nil
	})
}
