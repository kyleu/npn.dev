package transform

import (
	"fmt"
	"strings"

	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/npncore"
	"github.com/rbretecher/go-postman-collection"
	"logur.dev/logur"
)

type Postman struct {
}

var _ RequestTransformer = (*Postman)(nil)

func (x *Postman) Key() string {
	return "postman"
}

func (x *Postman) Description() string {
	return "TODO: postman"
}

func (x *Postman) TransformRequest(proto *request.Prototype, sess *session.Session, logger logur.Logger) (*Result, error) {
	pReq := toPostmanRequest(proto)
	out := npncore.ToJSON(pReq, logger)
	return &Result{Out: out}, nil
}

func (x *Postman) TransformCollection(coll *collection.Collection, requests request.Requests, sess *session.Session, logger logur.Logger) (*Result, error) {
	pColl := postman.CreateCollection(coll.TitleWithFallback(), coll.Description)
	for _, r := range requests {
		pColl.AddItem(postman.CreateItem(postman.Item{
			Name:        r.TitleWithFallback(),
			Description: r.Description,
			ID:          r.Key,
			Request:     toPostmanRequest(r.Prototype),
			Response:    nil,
		}))
	}
	out := npncore.ToJSON(pColl, logger)
	return &Result{Out: out}, nil
}

func toPostmanURL(p *request.Prototype) *postman.URL {
	return &postman.URL{
		Raw:       p.URLString(),
		Protocol:  p.Protocol.String(),
		Host:      []string{p.Domain},
		Path:      strings.Split(p.Path, "/"),
		Port:      fmt.Sprintf("%v", p.Port),
		Query:     nil,
		Hash:      p.Fragment,
		Variables: nil,
	}
}

func toPostmanRequest(p *request.Prototype) *postman.Request {
	return &postman.Request{
		URL:         toPostmanURL(p),
		Auth:        nil,
		Proxy:       nil,
		Certificate: nil,
		// Method:      ???,
		Description: nil,
		Header:      nil,
		Body:        nil,
	}
}

var txPostman = &Postman{}