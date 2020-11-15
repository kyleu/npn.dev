package socket

import (
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
)

// Collection
type collDetailsIn struct {
	Key        string                 `json:"key"`
	Collection *collection.Collection `json:"collection,omitempty"`
	Requests   request.Summaries      `json:"requests,omitempty"`
}

type addCollIn struct {
	Collections collection.Summaries `json:"collections"`
	Active      string               `json:"active"`
	Requests    request.Summaries    `json:"requests"`
}

type addURLOut struct {
	Coll string `json:"coll"`
	URL  string `json:"url"`
}

type addURLIn struct {
	Coll *collDetailsIn   `json:"coll"`
	Req  *request.Request `json:"req"`
}

type saveCollOut struct {
	OriginalKey string                 `json:"originalKey"`
	Coll        *collection.Collection `json:"coll"`
}

// Request
type reqDetailIn struct {
	Coll string           `json:"coll"`
	Req  *request.Request `json:"req"`
}

type reqDeleted struct {
	Req      string            `json:"req"`
	Coll     string            `json:"coll"`
	Requests request.Summaries `json:"requests"`
}

type getRequestOut struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
}

type saveRequestOut struct {
	Coll string           `json:"coll"`
	Orig string           `json:"orig"`
	Req  *request.Request `json:"req"`
}

type deleteRequestOut struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
}

type callOut struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req"`
	Proto *request.Prototype `json:"proto"`
}

// Transform
type transformOut struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req"`
	Fmt   string             `json:"fmt"`
	Proto *request.Prototype `json:"proto"`
}

type transformIn struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
	Fmt  string `json:"fmt"`
	Out  string `json:"out"`
}
