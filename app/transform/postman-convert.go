package transform

import (
	"bytes"

	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/rbretecher/go-postman-collection"
)

func PostmanImport(data []byte) (*postman.Collection, error) {
	coll, err := postman.ParseCollection(bytes.NewReader(data))
	return coll, err
}

func PostmanToFullCollection(pm *postman.Collection) (*collection.FullCollection, error) {
	coll := postmanToCollection(pm)
	reqs, sess, err := postmanToRequests(pm)
	if err != nil {
		return nil, err
	}

	ret := &collection.FullCollection{Coll: coll, Requests: reqs, Sess: sess}
	return ret, nil
}

func postmanToRequests(pm *postman.Collection) (request.Requests, *session.Session, error) {
	reqs := make(request.Requests, 0)
	sess := &session.Session{}

	for _, i := range pm.Items {
		newReqs, err := postmanItemToRequests(i, sess)
		if err != nil {
			return nil, nil, err
		}
		reqs = append(reqs, newReqs...)
	}

	return reqs, sess, nil
}

func postmanItemToRequests(pi *postman.Items, s *session.Session) (request.Requests, error) {
	ret := make(request.Requests, 0)
	if pi.Request != nil {
		ret = append(ret, parsePostmanRequest(pi.ID, pi.Name, pi.Description, pi.Request))
	}
	if pi.Variables != nil {
		for _, v := range pi.Variables {
			key := v.Name
			if len(key) == 0 {
				key = v.ID
			}
			if len(key) == 0 {
				key = v.Key
			}
			s.AddVariables(&session.Variable{Key: key, Value: v.Value})
		}
	}
	for _, i := range pi.Items {
		x, err := postmanItemToRequests(i, s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, x...)
	}
	return ret, nil
}

func parsePostmanRequest(id string, name string, desc string, r *postman.Request) *request.Request {
	proto := request.PrototypeFromString(r.URL.String())
	return &request.Request{Key: id, Title: name, Description: desc, Prototype: proto}
}

func postmanToCollection(i *postman.Collection) *collection.Collection {
	return &collection.Collection{Key: npncore.Slugify(i.Info.Name), Title: i.Info.Name, Description: i.Info.Description}
}
