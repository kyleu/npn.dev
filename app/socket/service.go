package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/app/search"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"

	"github.com/kyleu/libnpn/npnservice/user"

	"github.com/kyleu/npn/app/call"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/collection"
	"logur.dev/logur"
)

type Dependencies struct {
	User       user.Service
	Session    *session.Service
	Collection *collection.Service
	Request    *request.Service
	Caller     *call.Service
	Search     *search.Service
}

func NewService(deps *Dependencies, logger logur.Logger) *npnconnection.Service {
	return npnconnection.NewService(logger, onOpen, handler, onClose, deps)
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
	go sendCollections(s, c)
	go sendRequests(s, c)
	go sendSessions(s, c)
	return nil
}

func onClose(*npnconnection.Service, *npnconnection.Connection) error {
	return nil
}

func ctx(s *npnconnection.Service) *Dependencies {
	return s.Context.(*Dependencies)
}
