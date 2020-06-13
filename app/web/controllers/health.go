package controllers

import (
	"github.com/kyleu/npn/app/web"
	"github.com/kyleu/npn/app/web/act"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		_, _ = w.Write([]byte("OK"))
		return "", nil
	})
}
