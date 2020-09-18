package body

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"io"
	"io/ioutil"
)

func Parse(contentType string, contentLength int64, rd io.ReadCloser) (*Body, error) {
	defer func() { _ = rd.Close() }()

	if contentLength > 1024*1024 {
		cfg := &Large{Filename: "TODO", Length: contentLength}
		return &Body{Type: KeyLarge, Config: cfg}, nil
	}

	b, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, errors.Wrap(err, "can't read body")
	}

	if len(b) == 0 {
		return nil, nil
	}

	return detect(contentType, b), nil
}

func detect(contentType string, b []byte) *Body {
	switch contentType {
	case "application/json":
		cfg, err := tryParseJson(b)
		if err != nil {
			return detect("", b)
		}
		return &Body{Type: KeyJSON, Config: cfg}
	case "text/html":
		cfg := &HTML{Content: string(b)}
		return &Body{Type: KeyHTML, Config: cfg}
	default:
		return &Body{Type: KeyError, Config: &Error{Message: "unhandled content type [" + contentType + "]"}}
	}
}

func tryParseJson(b []byte) (Config, error) {
	var x interface{}
	err := npncore.FromJSON(b, &x)
	if err != nil {
		return nil, err
	}
	return &JSON{Msg: x, Length: int64(len(b))}, nil
}
