package controllers

import (
	"github.com/kyleu/npn/npnasset"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kyleu/npn/app/assets"
)

const assetBase = "web/assets"

func Favicon(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetBase, "/favicon.ico")
	npnasset.ZipResponse(w, r, data, hash, contentType, err)
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetBase, "/robots.txt")
	npnasset.ZipResponse(w, r, data, hash, contentType, err)
}

func Static(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(strings.TrimPrefix(r.URL.Path, "/assets"))
	if err == nil {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		data, hash, contentType, err := assets.Asset(assetBase, path)
		npnasset.ZipResponse(w, r, data, hash, contentType, err)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
