package body

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"strings"

	"emperror.dev/errors"
	"github.com/andybalholm/brotli"
	"github.com/kyleu/libnpn/npncore"
)

func Parse(path string, contentEncoding string, contentType string, charset string, contentLength int64, rd io.ReadCloser) (*Body, error) {
	defer func() { _ = rd.Close() }()

	if contentLength > 1024*1024 {
		return NewLarge("LARGE_FILENAME", contentType, contentLength), nil
	}

	var b []byte
	var err error
	contentEncoding, _ = npncore.SplitStringLast(contentEncoding, ';', true)
	switch contentEncoding {
	case "", "identity":
		b, err = ioutil.ReadAll(rd)
	case "gzip":
		b, err = gz(rd)
	case "deflate":
		b, err = deflate(rd)
	case "br", "brotli":
		b, err = br(rd)
	default:
		return nil, errors.New("unhandled encoding [" + contentEncoding + "]")
	}
	if err != nil {
		return nil, errors.Wrap(err, "can't read body")
	}

	if len(b) == 0 {
		return nil, nil
	}

	extension := path
	if strings.Contains(path, ".") {
		extension = path[strings.LastIndex(path, "."):]
	}

	return detect(extension, contentType, charset, b), nil
}

func gz(rd io.ReadCloser) ([]byte, error) {
	zr, _ := gzip.NewReader(rd)
	return ioutil.ReadAll(zr)
}

func deflate(rd io.ReadCloser) ([]byte, error) {
	var src []byte
	src, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, errors.Wrap(err, "error reading body")
	}
	dr := flate.NewReader(bytes.NewReader(src))
	var b []byte
	b, err = ioutil.ReadAll(dr)
	if err != nil {
		origErr := err
		zr, _ := zlib.NewReader(bytes.NewReader(src))
		b, err = ioutil.ReadAll(zr)
		if err != nil {
			err = origErr
		}
	}
	return b, err
}

func br(rd io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(brotli.NewReader(rd))
}

func detect(extension string, contentType string, charset string, b []byte) *Body {
	switch {
	case contentType == "application/json" || contentType == "text/javascript" || extension == "json" || extension == "js":
		return parseJSON(contentType, charset, b)
	case contentType == "text/html" || extension == "html":
		return parseHTML(b)
	case contentType == "text/xml" || contentType == "application/xml" || extension == "xml":
		return parseXML(b)
	case strings.HasPrefix(contentType, "image/"):
		return parseImage(contentType, b)
	default:
		return parseRaw(b)
	}
}

func parseRaw(b []byte) *Body {
	return NewRaw(b)
}
