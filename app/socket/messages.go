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
	ClientMessageGetSession    = "getSession"
	ClientMessageAddSession    = "addSession"
	ClientMessageSaveSession   = "saveSession"
	ClientMessageDeleteSession = "deleteSession"
)

const (
	ServerMessageLog = "log"

	// Collection
	ServerMessageCollections         = "collections"
	ServerMessageCollectionDetail    = "collectionDetail"
	ServerMessageCollectionNotFound  = "collectionNotFound"
	ServerMessageCollectionAdded     = "collectionAdded"
	ServerMessageCollectionUpdated   = "collectionUpdated"
	ServerMessageCollectionDeleted   = "collectionDeleted"
	ServerMessageCollectionTransform = "collectionTransform"

	// Request
	ServerMessageRequestDetail    = "requestDetail"
	ServerMessageRequestNotFound  = "requestNotFound"
	ServerMessageRequestAdded     = "requestAdded"
	ServerMessageRequestDeleted   = "requestDeleted"
	ServerMessageRequestTransform = "requestTransform"

	ServerMessageRequestStarted   = "requestStarted"
	ServerMessageRequestCompleted = "requestCompleted"

	// Session
	ServerMessageSessions         = "sessions"
	ServerMessageSessionAdded     = "sessionAdded"
	ServerMessageSessionDetail    = "sessionDetail"
	ServerMessageSessionDeleted   = "sessionDeleted"
	ServerMessageSessionNotFound  = "sessionNotFound"
	ServerMessageSessionTransform = "sessionTransform"
)
