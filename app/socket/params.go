package socket

import (
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
)

// Collection
type collDetailsOut struct {
	Key        string                 `json:"key"`
	Collection *collection.Collection `json:"collection,omitempty"`
	Requests   request.Summaries      `json:"requests,omitempty"`
}

type addCollOut struct {
	Collections collection.Summaries `json:"collections"`
	Active      string               `json:"active"`
	Requests    request.Summaries    `json:"requests"`
}

type addURLIn struct {
	Coll string `json:"coll"`
	URL  string `json:"url"`
}

type addURLOut struct {
	Coll *collDetailsOut  `json:"coll"`
	Req  *request.Request `json:"req"`
}

type saveCollOut struct {
	OriginalKey string                 `json:"originalKey"`
	Coll        *collection.Collection `json:"coll"`
}

type transformCollectionIn struct {
	Coll string `json:"coll"`
	Fmt  string `json:"fmt"`
}

// Request
type reqDetailOut struct {
	Coll    string           `json:"coll"`
	OrigKey string           `json:"origKey"`
	Req     *request.Request `json:"req"`
}

type reqDeletedOut struct {
	Req      string            `json:"req"`
	Coll     string            `json:"coll"`
	Requests request.Summaries `json:"requests"`
}

type getRequestIn struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
}

type saveRequestIn struct {
	Coll string           `json:"coll"`
	Orig string           `json:"orig"`
	Req  *request.Request `json:"req"`
}

type deleteRequestIn struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
}

type callIn struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req"`
	Sess  string             `json:"sess"`
	Proto *request.Prototype `json:"proto"`
}

type transformIn struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req,omitempty"`
	Sess  string             `json:"sess"`
	Fmt   string             `json:"fmt"`
	Proto *request.Prototype `json:"proto"`
}

type transformOut struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
	Fmt  string `json:"fmt"`
	Out  string `json:"out"`
}

// Session
type saveSessionIn struct {
	Orig string           `json:"orig"`
	Sess *session.Session `json:"sess"`
}

type addSessionOut struct {
	Sessions session.Summaries `json:"sessions"`
	Active   *session.Session  `json:"active"`
}
