package socket

import (
	"encoding/json"
	"fmt"
	"strings"

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

type addCollResult struct {
	Collections collection.Collections      `json:"collections"`
	Active      string                      `json:"active"`
	Requests    collection.RequestSummaries `json:"requests"`
}

type addURLInput struct {
	Coll string `json:"coll"`
	URL  string `json:"url"`
}

func handleCollectionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetCollection:
		return getCollDetails(s, c, param)
	case ClientMessageAddCollection:
		return addCollection(s, c, param)
	case ClientMessageAddRequestURL:
		return addRequestURL(s, c, param)
	default:
		return errors.New("unhandled collection command [" + cmd + "]")
	}
}

func addCollection(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	name := ""
	err := npncore.FromJSONStrict(param, &name)
	if err != nil {
		return errors.Wrap(err, "unable to parse input from URL")
	}
	key := npncore.Slugify(name)
	svcs := ctx(s)
	curr, _ := svcs.Collection.Load(key)
	if curr != nil {
		key += "-" + strings.ToLower(npncore.RandomString(4))
	}

	err = svcs.Collection.Save("", key, name, "")
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	newColls, _ := svcs.Collection.List()

	ret := &addCollResult{Collections: newColls, Active: key}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionAdded, ret)
	return s.WriteMessage(c.ID, msg)
}

func addRequestURL(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	p := &addURLInput{}
	err := npncore.FromJSONStrict(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse input from URL")
	}
	req, err := request.FromString("new", p.URL)
	if err != nil {
		return errors.Wrap(err, "unable to parse request from URL ["+p.URL+"]")
	}
	req.Key = npncore.Slugify(req.Prototype.Domain)

	svcs := ctx(s)
	curr, _ := svcs.Collection.LoadRequest(p.Coll, req.Key)
	if curr != nil {
		if req.Prototype != nil && len(req.Prototype.Path) > 0 {
			add := req.Prototype.Path
			if len(add) > 8 {
				add = add[0:8]
			}
			req.Key += "-" + npncore.Slugify(add)
		}
		curr, _ = svcs.Collection.LoadRequest(p.Coll, req.Key)
		if curr != nil {
			req.Key += "-" + strings.ToLower(npncore.RandomString(4))
		}
	}

	err = svcs.Collection.SaveRequest(p.Coll, "", req)
	if err != nil {
		return errors.Wrap(err, "unable to save request from URL ["+p.URL+"]")
	}

	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestAdded, &req)
	return s.WriteMessage(c.ID, msg)
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
