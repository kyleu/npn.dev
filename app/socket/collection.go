package socket

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
)

func handleCollectionMessage(s *npnconnection.Service, c *npnconnection.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetCollection:
		return getCollDetails(s, c, param)
	case ClientMessageAddCollection:
		return addCollection(s, c, param)
	case ClientMessageSaveCollection:
		return saveCollection(s, c, param)
	case ClientMessageDeleteCollection:
		return deleteCollection(s, c, param)
	case ClientMessageAddRequestURL:
		return addRequestURL(s, c, param)
	case ClientMessageTransform:
		return onTransformCollection(c, param, s)
	default:
		return errors.New("unhandled collection command [" + cmd + "]")
	}
}

func sendCollections(s *npnconnection.Service, c *npnconnection.Connection) {
	svcs := ctx(s)
	colls, err := svcs.Collection.Counts(&c.Profile.UserID)
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
	name, err := npncore.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	if len(name) == 0 {
		name = "new"
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

	newColls, _ := svcs.Collection.Counts(&c.Profile.UserID)

	ret := &addCollOut{Collections: newColls, Active: key}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionAdded, ret)
	return s.WriteMessage(c.ID, msg)
}

func saveCollection(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	p := &saveCollOut{}
	err := npncore.FromJSONStrict(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}

	svcs := ctx(s)
	err = svcs.Collection.Save(&c.Profile.UserID, p.OriginalKey, p.Coll.Key, p.Coll.Title, p.Coll.Description)
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+p.Coll.Key+"]")
	}

	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionUpdated, p.Coll)
	return s.WriteMessage(c.ID, msg)
}

func deleteCollection(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	key, err := npncore.FromJSONString(param)
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

func parseCollDetails(s *npnconnection.Service, userID *uuid.UUID, key string) (*collDetailsOut, error) {
	svcs := ctx(s)
	coll, err := svcs.Collection.Load(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	if coll == nil {
		return nil, nil
	}
	reqs, err := svcs.Request.ListRequests(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving requests for collection [%v]: %+v", key, err))
	}
	cd := &collDetailsOut{Key: key, Collection: coll, Requests: reqs}
	return cd, nil
}

func sendRequests(s *npnconnection.Service, c *npnconnection.Connection) {
	cd, _ := parseCollDetails(s, &c.Profile.UserID, "_")
	if cd == nil {
		msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionNotFound, &cd)
		_ = s.WriteMessage(c.ID, msg)
	}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionDetail, cd)
	_ = s.WriteMessage(c.ID, msg)
}

func getCollDetails(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	key, err := npncore.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing collection [%v]: %+v", key, err))
	}
	cd, err := parseCollDetails(s, &c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	if cd == nil {
		msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionNotFound, &cd)
		return s.WriteMessage(c.ID, msg)
	}
	msg := npnconnection.NewMessage(npncore.KeyCollection, ServerMessageCollectionDetail, cd)
	return s.WriteMessage(c.ID, msg)
}
