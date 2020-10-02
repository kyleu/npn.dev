package socket

import (
	"encoding/json"
	"fmt"

	"github.com/kyleu/npn/app/request"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func sendCollections(s *npnconnection.Service, connID uuid.UUID) {
	svcs := ctx(s)
	colls, err := svcs.Collection.List()
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error retrieving collections: %+v", err))
	}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollections, colls)
	err = s.WriteMessage(connID, msg)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

type collDetails struct {
	Collection *collection.Collection      `json:"collection"`
	Requests   collection.RequestSummaries `json:"requests"`
}

type addURLInput struct {
	Coll string `json:"coll"`
	URL  string `json:"url"`
}

func handleCollectionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetCollection:
		return getCollDetails(s, c, param)
	case ClientMessageAddURL:
		return addURL(s, c, param)
	default:
		return errors.New("unhandled collection command [" + cmd + "]")
	}
}

func addURL(s *npnconnection.Service, _ *npnconnection.Connection, param json.RawMessage) error {
	p := &addURLInput{}
	err := npncore.FromJSONStrict(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse input from URL")
	}
	req, err := request.FromString("new", p.URL)
	if err != nil {
		return errors.Wrap(err, "unable to parse request from URL ["+p.URL+"]")
	}
	req.Key = req.Prototype.Domain

	svcs := ctx(s)
	err = svcs.Collection.SaveRequest(p.Coll, "", req)
	if err != nil {
		return errors.Wrap(err, "unable to save request from URL ["+p.URL+"]")
	}

	return nil
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
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionDetail, &cd)
	return s.WriteMessage(c.ID, msg)
}
