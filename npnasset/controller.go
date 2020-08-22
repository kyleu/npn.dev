package npnasset

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kyleu/npn/npncontroller"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"github.com/kyleu/npn/npnasset/assets"
)

func VendorAsset(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(strings.TrimPrefix(r.URL.Path, "/vendor"))
	if err == nil {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		data, hash, contentType, err := assets.Asset("assets", path)
		ZipResponse(w, r, data, hash, contentType, err)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func ZipResponse(w http.ResponseWriter, r *http.Request, data []byte, hash string, contentType string, err error) {
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
