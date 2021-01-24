package socket

import (
	"encoding/json"


	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"

	"emperror.dev/errors"
)

func handleImportMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetImport:
		return getImport(s, c, param)
	default:
		return errors.New("unhandled import command [" + cmd + "]")
	}
}

func getImport(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	key := ""
	err := npncore.FromJSON(param, &key)
	if err != nil {
		return errors.Wrap(err, "unable to parse import key")
	}
	cfg, results, err := ctx(s).Import.Load(key)
	if err != nil {
		return errors.Wrap(err, "can't load import [" + key + "]")
	}
	ret := map[string]interface{}{"key": key, "cfg": cfg, "results": results}
	msg := npnconnection.NewMessage(npncore.KeyImport, ServerMessageImportResult, ret)
	return s.WriteMessage(c.ID, msg)
}
