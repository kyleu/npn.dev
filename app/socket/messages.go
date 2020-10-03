package socket

import "github.com/kyleu/npn/app/request"

const (
	ClientMessagePing    = "ping"
	ClientMessageConnect = "connect"

	// Collection
	ClientMessageGetCollection = "getCollection"
	ClientMessageAddURL        = "addURL"

	// Request
	ClientMessageGetRequest = "getRequest"
	ClientMessageCall       = "call"
	ClientMessageTransform  = "transform"
)

const (
	ServerMessagePong      = "pong"
	ServerMessageConnected = "connected"

	// Collection
	ServerMessageCollections      = "collections"
	ServerMessageCollectionDetail = "collectionDetail"

	// Request
	ServerMessageRequestDetail   = "requestDetail"
	ServerMessageCallResult      = "callResult"
	ServerMessageTransformResult = "transformResult"
)

type paramGetRequest struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
}

type paramCall struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req"`
	Proto *request.Prototype `json:"proto"`
}

type paramTransform struct {
	Coll  string             `json:"coll"`
	Req   string             `json:"req"`
	Fmt   string             `json:"fmt"`
	Proto *request.Prototype `json:"proto"`
}

type transformResponse struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
	Fmt  string `json:"fmt"`
	Out  string `json:"out"`
}
