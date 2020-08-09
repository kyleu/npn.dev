package npncore

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func GetUUIDFromString(s string) *uuid.UUID {
	var retID *uuid.UUID

	if len(s) > 0 {
		s, err := uuid.FromString(s)

		if err == nil {
			retID = &s
		}
	}

	return retID
}

func UUID() uuid.UUID {
	ret, err := uuid.NewV4()
	if err != nil {
		panic(errors.New("unable to create random UUID"))
	}

	return ret
}
