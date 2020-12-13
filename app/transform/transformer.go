package transform

import (
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Result struct {
	Key string `json:"key"`
	Out string `json:"out,omitempty"`
}

type Transformer interface {
	Key() string
	Title() string
	Description() string
}

type RequestTransformer interface {
	Transformer
	TransformRequest(proto *request.Prototype, sess *session.Session, logger logur.Logger) (*Result, error)
}

type RequestTransformers []RequestTransformer

func (t RequestTransformers) Get(s string) RequestTransformer {
	for _, x := range t {
		if x.Key() == s {
			return x
		}
	}
	return nil
}

var AllRequestTransformers = RequestTransformers{txCURL, txHTTP, txJSON, txOpenAPI, txPostman}

type CollectionTransformer interface {
	Transformer
	TransformCollection(coll *collection.Collection, requests request.Requests, sess *session.Session, logger logur.Logger) (*Result, error)
}

type CollectionTransformers []CollectionTransformer

func (t CollectionTransformers) Get(s string) CollectionTransformer {
	for _, x := range t {
		if x.Key() == s {
			return x
		}
	}
	return nil
}

var AllCollectionTransformers = CollectionTransformers{txJSON, txOpenAPI, txPostman}

func TransformSession(sess *session.Session, logger logur.Logger) (*Result, error) {
	return &Result{Key: sess.Key, Out: npncore.ToJSON(sess, logger)}, nil
}
