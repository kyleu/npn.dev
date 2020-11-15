package socket

const (
	ClientMessageTestbed     = "testbed"
	ClientMessageSaveProfile = "saveProfile"

	// Collection
	ClientMessageGetCollection    = "getCollection"
	ClientMessageAddCollection    = "addCollection"
	ClientMessageSaveCollection   = "saveCollection"
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
	// Collection
	ServerMessageCollections        = "collections"
	ServerMessageCollectionDetail   = "collectionDetail"
	ServerMessageCollectionNotFound = "collectionNotFound"
	ServerMessageCollectionAdded    = "collectionAdded"
	ServerMessageCollectionUpdated  = "collectionUpdated"
	ServerMessageCollectionDeleted  = "collectionDeleted"

	// Request
	ServerMessageRequestDetail   = "requestDetail"
	ServerMessageRequestNotFound = "requestNotFound"
	ServerMessageRequestAdded    = "requestAdded"
	ServerMessageRequestDeleted  = "requestDeleted"
	ServerMessageCallResult      = "callResult"
	ServerMessageTransformResult = "transformResult"
)
