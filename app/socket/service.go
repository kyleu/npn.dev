package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type services struct {
	Collection *collection.Service
}

func NewService(collection *collection.Service, logger logur.Logger) *npnconnection.Service {
	ctx := &services{
		Collection: collection,
	}
	return npnconnection.NewService(logger, handler, ctx)
}

func handler(s *npnconnection.Service, c *npnconnection.Connection, svc string, cmd string, param json.RawMessage) error {
	var err error
	switch svc {
	case npncore.KeySystem:
		err = handleSystemMessage(s, c, cmd, param)
	default:
		err = errors.New(npncore.IDErrorString(npncore.KeyService, svc))
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

func ctx(s *npnconnection.Service) *services {
	return s.Context.(*services)
}
