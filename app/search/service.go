package search

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"logur.dev/logur"
)

type Service struct {
	coll   *collection.Service
	req    *request.Service
	logger logur.Logger
}

func NewService(coll *collection.Service, req *request.Service, logger logur.Logger) *Service {
	return &Service{coll: coll, req: req, logger: logger}
}

func (s Service) Run(p *Params, userID *uuid.UUID, role string) (Results, error) {
	colls, err := s.coll.List(userID)
	if err != nil {
		return nil, err
	}

	ret := make(Results, 0)

	matchedColls := searchCollections(colls, p.Q)
	ret = append(ret, matchedColls...)

	matchedReqs, err := searchRequests(userID, colls, p.Q, s.req)
	if err != nil {
		return nil, err
	}
	ret = append(ret, matchedReqs...)

	return ret, nil
}

func searchCollections(colls collection.Collections, q string) Results {
	ret := make(Results, 0, len(colls))
	for _, coll := range colls {
		matched, prelude, loc, postlude := coll.Matches(q)
		if matched {
			ret = append(ret, &Result{
				T:        "collection",
				Key:      coll.Key,
				Prelude:  prelude,
				Loc:      loc,
				Postlude: postlude,
			})
		}
	}
	return ret
}

func searchRequests(userID *uuid.UUID, colls collection.Collections, q string, reqSvc *request.Service) (Results, error) {
	ret := make(Results, 0)
	for _, coll := range colls {
		reqs, err := reqSvc.List(userID, coll.Key)
		if err != nil {
			return nil, err
		}
		for _, req := range reqs {
			matched, prelude, loc, postlude := req.Matches(q)
			if matched {
				ret = append(ret, &Result{
					T:        "request",
					Key:      coll.Key + "/" + req.Key,
					Prelude:  prelude,
					Loc:      loc,
					Postlude: postlude,
				})
			}
		}
	}
	return ret, nil
}
