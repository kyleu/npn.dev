package socket

import (
	"github.com/kyleu/npn/npnuser"

	"emperror.dev/errors"
)

type ConnectedResponse struct {
	Profile *npnuser.Profile `json:"profile"`
}

func handleSystemMessage(cmd string) error {
	return errors.New("unhandled app command [" + cmd + "]")
}
