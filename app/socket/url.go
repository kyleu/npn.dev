package socket

import (
	"encoding/json"
	"strings"

	"github.com/kyleu/npn/app/request"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
)

func addRequestURL(s *npnconnection.Service, c *npnconnection.Connection, param json.RawMessage) error {
	p := &addURLOut{}
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
	curr, _ := svcs.Request.LoadRequest(&c.Profile.UserID, p.Coll, req.Key)
	if curr != nil {
		clean(req)
		curr, _ = svcs.Request.LoadRequest(&c.Profile.UserID, p.Coll, req.Key)
		if curr != nil {
			req.Key += "-" + strings.ToLower(npncore.RandomString(4))
		}
	}

	err = svcs.Request.SaveRequest(&c.Profile.UserID, p.Coll, "", req)
	if err != nil {
		return errors.Wrap(err, "unable to save request from URL ["+p.URL+"]")
	}

	coll, err := parseCollDetails(s, &c.Profile.UserID, p.Coll)
	if err != nil {
		return err
	}

	out := &addURLIn{
		Coll: coll,
		Req:  req,
	}
	msg := npnconnection.NewMessage(npncore.KeyRequest, ServerMessageRequestAdded, out)
	return s.WriteMessage(c.ID, msg)
}

func clean(req *request.Request) {
	if len(req.Title) == 0 {
		req.Title = req.Key
	}
	if req.Prototype != nil && len(req.Prototype.Path) > 0 {
		add := req.Prototype.Path
		if len(add) > 8 {
			add = add[0:8]
		}
		req.Key += "-" + npncore.Slugify(add)
	}
}
