package socket

import (
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func sendCollections(s *npnconnection.Service, connID uuid.UUID) {
	svcs := ctx(s)
	colls, err := svcs.Collection.List()
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error retrieving collections: %+v", err))
	}
	msg := npnconnection.NewMessage("collection", "collections", colls)
	err = s.WriteMessage(connID, msg)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

type collDetails struct {
	Collection *collection.Collection `json:"collection"`
	Requests   request.Requests       `json:"requests"`
}

func handleCollectionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case "getCollection":
		return getCollDetails(s, c, param)
	default:
		return errors.New("unhandled app command [" + cmd + "]")
	}
}

func getCollDetails(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	svcs := ctx(s)
	key := ""
	err := npncore.FromJSON(param, &key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing collection [%v]: %+v", key, err))
	}
	coll, err := svcs.Collection.Load(key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	reqs, err := svcs.Collection.ListRequests(key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error retrieving requests for collection [%v]: %+v", key, err))
	}
	cd := collDetails{Collection: coll, Requests: reqs}
	msg := npnconnection.NewMessage("collection", "collectionDetail", &cd)
	return s.WriteMessage(c.ID, msg)
}
