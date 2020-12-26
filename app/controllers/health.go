package controllers

import (
	"net/http"

	"github.com/kyleu/libnpn/npncore"

	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npnweb"
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
