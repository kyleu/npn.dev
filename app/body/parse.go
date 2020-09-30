package body

import (
	"compress/flate"
	"compress/gzip"
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"io"
	"io/ioutil"
	"strings"
)

func Parse(contentEncoding string, contentType string, charset string, contentLength int64, rd io.ReadCloser) (*Body, error) {
	defer func() { _ = rd.Close() }()

	if contentLength > 1024*1024 {
		cfg := &Large{Filename: "TODO", Length: contentLength}
		return &Body{Type: KeyLarge, Config: cfg}, nil
	}

	var b []byte
	var err error
	contentEncoding, _ = npncore.SplitStringLast(contentEncoding, ';', true)
	switch contentEncoding {
	case "", "identity":
		b, err = ioutil.ReadAll(rd)
	case "deflate":
		dr := flate.NewReader(rd)
		b, err = ioutil.ReadAll(dr)
	case "gzip":
		zr, _ := gzip.NewReader(rd)
		b, err = ioutil.ReadAll(zr)
	default:
		return nil, errors.New("unhandled encoding [" + contentEncoding + "]")
	}
	if err != nil {
		return nil, errors.Wrap(err, "can't read body")
	}

	if len(b) == 0 {
		return nil, nil
	}

	return detect(contentType, charset, b), nil
}

func detect(contentType string, charset string, b []byte) *Body {
	switch {
	case contentType == "application/json":
		return parseJSON(contentType, charset, b)
	case contentType == "text/html":
		return parseHTML(b)
	case strings.HasPrefix(contentType, "image/"):
		return parseImage(contentType, b)
	default:
		return parseRaw(b)
	}
}

func parseRaw(b []byte) *Body {
	return NewRaw(b)
}
