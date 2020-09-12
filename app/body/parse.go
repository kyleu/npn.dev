package body

import (
	"emperror.dev/errors"
	"io"
)

func Parse(contentType string, reader io.ReadCloser) (*Body, error) {
	err := reader.Close()
	if err != nil {
		return nil, errors.Wrap(err, "unable to close body")
	}

	return nil, nil
}
