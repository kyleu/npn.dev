package socket

const (
	ClientMessageTestbed     = "testbed"
	ClientMessageSaveProfile = "saveProfile"

	ClientMessageSearch    = "search"
	ClientMessageRunURL    = "runURL"
	ClientMessageTransform = "transform"

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

	// Session
	ClientMessageGetSession    = "getSession"
	ClientMessageAddSession    = "addSession"
	ClientMessageSaveSession   = "saveSession"
	ClientMessageDeleteSession = "deleteSession"

	// Imports
	ClientMessageGetImport = "getImport"
)

const (
	ServerMessageSearchResults = "searchResults"

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

	// Imports
	ServerMessageImportResult = "importResult"
)
