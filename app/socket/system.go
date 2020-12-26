package socket

import (
	"encoding/json"

	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnuser"

	"emperror.dev/errors"
)

func handleSystemMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageTestbed:
		return testbed(param, s)
	case ClientMessageSaveProfile:
		return saveProfile(s, c, param)
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}

func saveProfile(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	p := &npnuser.Profile{}
	err := npncore.FromJSON(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse profile")
	}
	p.UserID = c.Profile.UserID
	p.Role = c.Profile.Role

	svcs := ctx(s)

	_, err = svcs.User.SaveProfile(p.ToUserProfile())
	return err
}
