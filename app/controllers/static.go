package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kyleu/npn/npncontroller"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/assets"
)

func Favicon(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset("web/assets", "/favicon.ico")
	zipResponse(w, r, data, hash, contentType, err)
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset("web/assets", "/robots.txt")
	zipResponse(w, r, data, hash, contentType, err)
}

func Static(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(strings.TrimPrefix(r.URL.Path, "/assets"))
	if err == nil {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		data, hash, contentType, err := assets.Asset("web/assets", path)
		zipResponse(w, r, data, hash, contentType, err)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func zipResponse(w http.ResponseWriter, r *http.Request, data []byte, hash string, contentType string, err error) {
	if err == nil {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", contentType)
		// w.Header().Add("Cache-Control", "public, max-age=31536000")
		w.Header().Add("ETag", hash)
		if r.Header.Get("If-None-Match") == hash {
			w.WriteHeader(http.StatusNotModified)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(data)
			emperror.Panic(errors.Wrap(err, "unable to write to response"))
		}
	} else {
		npncontroller.NotFound(w, r)
	}
}
