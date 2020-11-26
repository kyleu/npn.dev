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
	ClientMessageRunURL        = "runURL"
	ClientMessageGetRequest    = "getRequest"
	ClientMessageSaveRequest   = "saveRequest"
	ClientMessageDeleteRequest = "deleteRequest"
	ClientMessageCall          = "call"
	ClientMessageTransform     = "transform"

	// Session
	ClientMessageGetSession = "getSession"
	ClientMessageSaveSession = "saveSession"
	ClientMessageDeleteSession = "deleteSession"
)

const (
	// Session
	ServerMessageSessions = "sessions"

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

	// Session
	ServerMessageSessionDetail   = "sessionDetail"
	ServerMessageSessionNotFound = "sessionNotFound"
)
