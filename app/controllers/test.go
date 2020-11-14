package controllers

import (
	"net/http"

	"github.com/mccutchen/go-httpbin/httpbin"
)

var testSvc *httpbin.HTTPBin

func TestCall(w http.ResponseWriter, r *http.Request) {
	if testSvc == nil {
		testSvc = httpbin.NewHTTPBin()
	}
	testSvc.Handler().ServeHTTP(w, r)
}
