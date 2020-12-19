package npncore

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

// Returns the UUID represented by the provided string. Only use with trusted input, this swallows errors
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

// Returns a random V4 UUID, panicing if an error happens, which it never does
func UUID() uuid.UUID {
	ret, err := uuid.NewV4()
	if err != nil {
		panic(errors.New("unable to create random UUID"))
	}

	return ret
}
