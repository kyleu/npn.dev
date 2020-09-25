package body

import (
	"compress/gzip"
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"io"
	"io/ioutil"
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
	case "gzip", "deflate":
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
	switch contentType {
	case "application/json":
		cfg, err := tryParseJSON(b)
		if err != nil {
			return detect("", charset, b)
		}
		return &Body{Type: KeyJSON, Config: cfg}
	case "text/html":
		cfg := &HTML{Content: string(b)}
		return &Body{Type: KeyHTML, Config: cfg}
	default:
		return &Body{Type: KeyError, Config: &Error{Message: "unhandled content type [" + contentType + "]"}}
	}
}

func tryParseJSON(b []byte) (Config, error) {
	var x interface{}
	err := npncore.FromJSON(b, &x)
	if err != nil {
		return nil, err
	}
	return &JSON{Msg: x}, nil
}
