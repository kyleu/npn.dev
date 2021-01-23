package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/app/search"

	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnuser"

	"emperror.dev/errors"
)

func handleSystemMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageTestbed:
		return testbed(param, s)
	case ClientMessageSearch:
		return onSearch(s, c, param)
	case ClientMessageSaveProfile:
		return saveProfile(s, c, param)
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}

func onSearch(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	sp := &search.Params{}
	err := npncore.FromJSON(param, sp)
	if err != nil {
		return errors.Wrap(err, "unable to parse search params")
	}
	results, err := ctx(s).Search.Run(sp, &c.Profile.UserID, c.Profile.Role)
	msg := npnconnection.NewMessage(npncore.KeySystem, ServerMessageSearchResults, results)
	return s.WriteMessage(c.ID, msg)
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
