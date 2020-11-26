package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"

	"github.com/kyleu/npn/npnservice/user"

	"github.com/kyleu/npn/app/call"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type services struct {
	User       user.Service
	Session    *session.Service
	Collection *collection.Service
	Request    *request.Service
	Caller     *call.Service
}

func NewService(userSvc user.Service, sessSvc *session.Service, collectionSvc *collection.Service, requestSvc *request.Service, callSvc *call.Service, logger logur.Logger) *npnconnection.Service {
	ctx := &services{
		User:       userSvc,
		Session:    sessSvc,
		Collection: collectionSvc,
		Request:    requestSvc,
		Caller:     callSvc,
	}
	return npnconnection.NewService(logger, onOpen, handler, onClose, ctx)
}

func handler(s *npnconnection.Service, c *npnconnection.Connection, svc string, cmd string, param json.RawMessage) error {
	var err error
	switch svc {
	case npncore.KeySystem:
		err = handleSystemMessage(s, c, cmd, param)
	case npncore.KeyCollection:
		err = handleCollectionMessage(s, c, cmd, param)
	case npncore.KeyRequest:
		err = handleRequestMessage(s, c, cmd, param)
	case npncore.KeySession:
		err = handleSessionMessage(s, c, cmd, param)
	default:
		err = errors.New(npncore.IDErrorString(npncore.KeyService, svc))
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

func onOpen(s *npnconnection.Service, c *npnconnection.Connection) error {
	go sendSessions(s, c)
	go sendCollections(s, c)
	return nil
}

func onClose(*npnconnection.Service, *npnconnection.Connection) error {
	return nil
}

func ctx(s *npnconnection.Service) *services {
	return s.Context.(*services)
}
