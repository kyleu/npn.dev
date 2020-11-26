package socket

import (
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func handleSessionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	var err error

	switch cmd {
	case ClientMessageGetSession:
		err = onGetSession(c, param, s)
	case ClientMessageSaveSession:
		err = onSaveSession(c, param, s)
	case ClientMessageDeleteSession:
		err = onDeleteSession(c, param, s)
	default:
		err = errors.New("invalid session command [" + cmd + "]")
	}

	return err
}

func sendSessions(s *npnconnection.Service, c *npnconnection.Connection) {
	svcs := ctx(s)
	sessions, err := svcs.Session.Counts(&c.Profile.UserID)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error retrieving sessions: %+v", err))
	}
	msg := npnconnection.NewMessage(npncore.KeySession, ServerMessageSessions, sessions)
	err = s.WriteMessage(c.ID, msg)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

func onGetSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	var key string
	err := npncore.FromJSON(param, &key)
	if err != nil {
		return errors.Wrap(err, "unable to read session key")
	}

	svcs := ctx(s)
	session, err := svcs.Session.Load(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to load session ["+key+"]")
	}
	var msg *npnconnection.Message
	if session == nil {
		msg = npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionNotFound, key)
	} else {
		msg = npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionDetail, session)
	}
	return s.WriteMessage(c.ID, msg)
}

func onSaveSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	msg := npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionNotFound, "TODO")
	return s.WriteMessage(c.ID, msg)
}

func onDeleteSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	msg := npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionNotFound, "TODO")
	return s.WriteMessage(c.ID, msg)
}
