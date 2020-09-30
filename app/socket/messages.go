package socket

const (
	ClientMessagePing    = "ping"
	ClientMessageConnect = "connect"

	// Collection
	ClientMessageGetCollection = "getCollection"
	ClientMessageAddURL        = "addURL"

	// Request
	ClientMessageGetRequest = "getRequest"
	ClientMessageCall       = "call"
)

const (
	ServerMessagePong = "pong"
	ServerMessageConnected = "connected"

	// Collection
	ServerMessageCollections = "collections"
	ServerMessageCollectionDetail = "collectionDetail"

	// Request
	ServerMessageRequestDetail = "requestDetail"
	ServerMessageCallResult = "callResult"
)
