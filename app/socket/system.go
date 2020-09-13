package socket

import (
	"encoding/json"
	"fmt"
	"github.com/kyleu/npn/npnuser"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

type connected struct {
	Profile npnuser.Profile `json:"profile"`
}

func handleSystemMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case "connect":
		param := connected{Profile: c.Profile}
		msg := npnconnection.NewMessage(npncore.KeySystem, "connected", param)
		err := s.WriteMessage(c.ID, msg)
		if err != nil {
			return errors.Wrap(err, "unable to write to socket")
		}
		go func() {
			svcs := ctx(s)
			colls, err := svcs.Collection.List()
			if err != nil {
				s.Logger.Warn(fmt.Sprintf("error retrieving collections: %+v", err))
			}
			msg := npnconnection.NewMessage("collection", "collections", colls)
			err = s.WriteMessage(c.ID, msg)
			if err != nil {
				s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
			}
		}()
		return nil
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}
