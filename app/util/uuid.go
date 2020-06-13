package util

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func UUID() uuid.UUID {
	ret, err := uuid.NewV4()
	if err != nil {
		panic(errors.New("unable to create random UUID"))
	}

	return ret
}
