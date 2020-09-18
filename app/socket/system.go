package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/npnuser"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
)

type connected struct {
	Profile npnuser.Profile `json:"profile"`
}

func handleSystemMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}
