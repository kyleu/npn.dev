package socket

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kyleu/npn/app/transform"

	"github.com/kyleu/npn/app/session"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func handleSessionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	var err error

	switch cmd {
	case ClientMessageGetSession:
		err = onGetSession(c, param, s)
	case ClientMessageAddSession:
		err = onAddSession(c, param, s)
	case ClientMessageSaveSession:
		err = onSaveSession(c, param, s)
	case ClientMessageDeleteSession:
		err = onDeleteSession(c, param, s)
	case ClientMessageTransform:
		return onTransformSession(c, param, s)
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
	key, err := npncore.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to read session key")
	}

	svcs := ctx(s)
	sess, err := svcs.Session.Load(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to load session ["+key+"]")
	}
	var msg *npnconnection.Message
	if sess == nil {
		msg = npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionNotFound, key)
	} else {
		msg = npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionDetail, sess)
	}
	return s.WriteMessage(c.ID, msg)
}

func onAddSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	name, err := npncore.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse session name")
	}
	if len(name) == 0 {
		name = "new"
	}
	key := npncore.Slugify(name)
	svcs := ctx(s)
	curr, _ := svcs.Session.Load(&c.Profile.UserID, key)
	if curr != nil {
		key += "-" + strings.ToLower(npncore.RandomString(4))
	}

	sess := &session.Session{Key: key, Title: name}
	err = svcs.Session.Save(&c.Profile.UserID, "", sess)
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	sessCounts, _ := svcs.Session.Counts(&c.Profile.UserID)

	ret := &addSessionOut{Sessions: sessCounts, Active: sess}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageSessionAdded, ret)
	return s.WriteMessage(c.ID, msg)
}

func onSaveSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	svc := ctx(s)
	frm := &saveSessionIn{}
	err := npncore.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveSession param")
	}
	frm.Sess = frm.Sess.Minify()
	err = svc.Session.Save(&c.Profile.UserID, frm.Orig, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't save session ["+frm.Sess.Key+"]")
	}
	msg := npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionDetail, frm.Sess.Normalize(frm.Sess.Key))
	return s.WriteMessage(c.ID, msg)
}

func onDeleteSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	key, err := npncore.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	svcs := ctx(s)
	err = svcs.Session.Delete(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to delete session with key ["+key+"]")
	}

	msg := npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionDeleted, key)
	return s.WriteMessage(c.ID, msg)
}

func onTransformSession(c *npnconnection.Connection, param json.RawMessage, s *npnconnection.Service) error {
	frm := ""
	err := npncore.FromJSONStrict(param, &frm)
	if err != nil {
		return errors.Wrap(err, "can't load session transform param")
	}

	sess, err := ctx(s).Session.Load(&c.Profile.UserID, frm)
	if err != nil {
		return errors.Wrap(err, "can't load session transform ["+frm+"]")
	}

	rsp, err := transform.Session(sess, s.Logger)
	if err != nil {
		return errors.Wrap(err, "can't load transform session")
	}

	msg := npnconnection.NewMessage(npncore.KeySession, ServerMessageSessionTransform, rsp)
	return s.WriteMessage(c.ID, msg)
}
