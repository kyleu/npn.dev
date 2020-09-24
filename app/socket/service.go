package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/app/call"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type services struct {
	Collection *collection.Service
	Caller     *call.Service
}

func NewService(collectionSvc *collection.Service, callSvc *call.Service, logger logur.Logger) *npnconnection.Service {
	ctx := &services{
		Collection: collectionSvc,
		Caller: callSvc,
	}
	return npnconnection.NewService(logger, onOpen, handler, onClose, ctx)
}

func handler(s *npnconnection.Service, c *npnconnection.Connection, svc string, cmd string, param json.RawMessage) error {
	var err error
	switch svc {
	case "collection":
		err = handleCollectionMessage(s, c, cmd, param)
	case npncore.KeySystem:
		err = handleSystemMessage(cmd)
	case npncore.KeyRequest:
		err = handleRequestMessage(s, c, cmd, param)
	default:
		err = errors.New(npncore.IDErrorString(npncore.KeyService, svc))
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

func onOpen(s *npnconnection.Service, c *npnconnection.Connection) error {
	p := connected{Profile: c.Profile}
	msg := npnconnection.NewMessage(npncore.KeySystem, "connected", p)
	err := s.WriteMessage(c.ID, msg)
	if err != nil {
		return errors.Wrap(err, "unable to write to socket")
	}
	go sendCollections(s, c.ID)
	return nil
}

func onClose(s *npnconnection.Service, c *npnconnection.Connection) error {
	return nil
}

func ctx(s *npnconnection.Service) *services {
	return s.Context.(*services)
}
