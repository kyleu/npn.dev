package socket

import (
	"fmt"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

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
