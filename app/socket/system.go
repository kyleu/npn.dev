package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/npnuser"

	"emperror.dev/errors"
)

type ConnectedResponse struct {
	Profile *npnuser.Profile `json:"profile"`
}

func handleSystemMessage(cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageTestbed:
		return testbed(param)
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}
