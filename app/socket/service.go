package socket

import (
	"encoding/json"

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
	case npncore.AppName:
		println("Hello!")
	default:
		err = errors.New(npncore.IDErrorString(npncore.KeyService, svc))
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

// func ctx(s *npnconnection.Service) *services {
// 	return s.Context.(*services)
// }
