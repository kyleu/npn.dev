package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/npnuser"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type services struct {
}

func NewService(logger logur.Logger) *npnconnection.Service {
	ctx := &services{}
	return npnconnection.NewService(logger, handler, ctx)
}

func handler(s *npnconnection.Service, c *npnconnection.Connection, svc string, cmd string, param json.RawMessage) error {
	var err error
	switch svc {
	case npncore.KeySystem:
		err = handleKeySystemMessage(s, c, cmd, param)
	default:
		err = errors.New(npncore.IDErrorString(npncore.KeyService, svc))
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

type connected struct {
	Profile npnuser.Profile `json:"profile"`
}

func handleKeySystemMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case "connect":
		param := connected{Profile: c.Profile}
		msg := npnconnection.NewMessage(npncore.KeySystem, "connected", param)
		return s.WriteMessage(c.ID, msg)
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}

func ctx(s *npnconnection.Service) *services {
	return s.Context.(*services)
}
