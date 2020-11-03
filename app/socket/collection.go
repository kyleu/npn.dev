package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"

	"github.com/kyleu/npn/app/request"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

type collDetails struct {
	Key        string                      `json:"key"`
	Collection *collection.Collection      `json:"collection,omitempty"`
	Requests   collection.RequestSummaries `json:"requests,omitempty"`
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

type addURLOutput struct {
	Coll *collDetails `json:"coll"`
	Req  *request.Request `json:"req"`
}

func handleCollectionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetCollection:
		return getCollDetails(s, c, param)
	case ClientMessageAddCollection:
		return addCollection(s, c, param)
	case ClientMessageDeleteCollection:
		return deleteCollection(s, c, param)
	case ClientMessageAddRequestURL:
		return addRequestURL(s, c, param)
	default:
		return errors.New("unhandled collection command [" + cmd + "]")
	}
}

func sendCollections(s *npnconnection.Service, c *npnconnection.Connection) {
	svcs := ctx(s)
	colls, err := svcs.Collection.List(&c.Profile.UserID)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error retrieving collections: %+v", err))
	}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollections, colls)
	err = s.WriteMessage(c.ID, msg)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
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
	curr, _ := svcs.Collection.Load(&c.Profile.UserID, key)
	if curr != nil {
		key += "-" + strings.ToLower(npncore.RandomString(4))
	}

	err = svcs.Collection.Save(&c.Profile.UserID, "", key, name, "")
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	newColls, _ := svcs.Collection.List(&c.Profile.UserID)

	ret := &addCollResult{Collections: newColls, Active: key}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionAdded, ret)
	return s.WriteMessage(c.ID, msg)
}

func deleteCollection(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	key := ""
	err := npncore.FromJSONStrict(param, &key)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	svcs := ctx(s)
	err = svcs.Collection.Delete(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to delete collection with key ["+key+"]")
	}

	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionDeleted, key)
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
	curr, _ := svcs.Collection.LoadRequest(&c.Profile.UserID, p.Coll, req.Key)
	if curr != nil {
		if req.Prototype != nil && len(req.Prototype.Path) > 0 {
			add := req.Prototype.Path
			if len(add) > 8 {
				add = add[0:8]
			}
			req.Key += "-" + npncore.Slugify(add)
		}
		curr, _ = svcs.Collection.LoadRequest(&c.Profile.UserID, p.Coll, req.Key)
		if curr != nil {
			req.Key += "-" + strings.ToLower(npncore.RandomString(4))
		}
	}

	err = svcs.Collection.SaveRequest(&c.Profile.UserID, p.Coll, "", req)
	if err != nil {
		return errors.Wrap(err, "unable to save request from URL ["+p.URL+"]")
	}

	coll, err := parseCollDetails(s, &c.Profile.UserID, p.Coll)
	if err != nil {
		return err
	}

	out := &addURLOutput{
		Coll: coll,
		Req:  req,
	}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestAdded, out)
	return s.WriteMessage(c.ID, msg)
}

func parseCollDetails(s *npnconnection.Service, userID *uuid.UUID, key string) (*collDetails, error) {
	svcs := ctx(s)
	coll, err := svcs.Collection.Load(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	reqs, err := svcs.Collection.ListRequests(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving requests for collection [%v]: %+v", key, err))
	}
	cd := &collDetails{Key: key, Collection: coll, Requests: reqs}
	return cd, nil
}

func getCollDetails(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	key := ""
	err := npncore.FromJSON(param, &key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing collection [%v]: %+v", key, err))
	}
	cd, err := parseCollDetails(s, &c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionDetail, &cd)
	return s.WriteMessage(c.ID, msg)
}
