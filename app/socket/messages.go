package socket

import "github.com/kyleu/npn/app/request"

const (
	ClientMessagePing    = "ping"
	ClientMessageConnect = "connect"

	// Collection
	ClientMessageGetCollection    = "getCollection"
	ClientMessageAddCollection    = "addCollection"
	ClientMessageDeleteCollection = "deleteCollection"
	ClientMessageAddRequestURL    = "addRequestURL"

	// Request
	ClientMessageGetRequest    = "getRequest"
	ClientMessageSaveRequest   = "saveRequest"
	ClientMessageDeleteRequest = "deleteRequest"
	ClientMessageCall          = "call"
	ClientMessageTransform     = "transform"
)

const (
	ServerMessagePong = "pong"

	// Collection
	ServerMessageCollections       = "collections"
	ServerMessageCollectionDetail  = "collectionDetail"
	ServerMessageCollectionAdded   = "collectionAdded"
	ServerMessageCollectionDeleted = "collectionDeleted"

	// Request
	ServerMessageRequestDetail   = "requestDetail"
	ServerMessageRequestAdded    = "requestAdded"
	ServerMessageRequestDeleted  = "requestDeleted"
	ServerMessageCallResult      = "callResult"
	ServerMessageTransformResult = "transformResult"
)

type paramGetRequest struct {
	Coll string `json:"coll"`
	Req  string `json:"req"`
}

type paramSaveRequest struct {
	Coll string           `json:"coll"`
	Orig string           `json:"orig"`
	Req  *request.Request `json:"req"`
}

type paramDeleteRequest struct {
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
