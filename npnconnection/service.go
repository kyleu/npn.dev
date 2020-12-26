package npnconnection

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/kyleu/npn/npnuser"

	"github.com/kyleu/npn/npncore"

	"github.com/gofrs/uuid"
	"logur.dev/logur"
)

// Function used to handle incoming messages
type Handler func(s *Service, conn *Connection, svc string, cmd string, param json.RawMessage) error
// Function used to handle incoming connections
type ConnectEvent func(s *Service, conn *Connection) error

// Manages all Connection objects
type Service struct {
	connections   map[uuid.UUID]*Connection
	connectionsMu sync.Mutex
	channels      map[Channel][]uuid.UUID
	channelsMu    sync.Mutex
	Logger        logur.Logger
	onOpen        ConnectEvent
	handler       Handler
	onClose       ConnectEvent
	wasmCallback  func(string)
	Context       interface{}
}

// Creates a new service with the provided handler functions
func NewService(logger logur.Logger, onOpen ConnectEvent, handler Handler, onClose ConnectEvent, ctx interface{}) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeySocket})
	return &Service{
		connections: make(map[uuid.UUID]*Connection),
		channels:    make(map[Channel][]uuid.UUID),
		Logger:      logger,
		handler:     handler,
		onOpen:      onOpen,
		Context:     ctx,
	}
}

var systemID = uuid.FromStringOrNil("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF")
var systemStatus = &Status{ID: systemID, UserID: systemID, Username: "System Broadcast", ChannelSvc: npncore.KeySystem, ChannelID: &systemID}

// Used by userless WASM messages
var WASMID = uuid.FromStringOrNil("CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC")
// Used by userless WASM messages
var WASMProfile = npnuser.NewUserProfile(WASMID, "WebAssembly Client").ToProfile()
var wasmStatus = &Status{ID: WASMID, UserID: WASMID, Username: "WebAssembly Client", ChannelSvc: npncore.KeySystem, ChannelID: &systemID}
var wasmConnection = &Connection{ID: WASMID, Profile: WASMProfile}

// Returns an array of Connection statuses based on the parameters
func (s *Service) List(params *npncore.Params) Statuses {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyConnection, params)
	ret := make(Statuses, 0)
	ret = append(ret, systemStatus)
	var idx = 0
	for _, conn := range s.connections {
		if idx >= params.Offset && (params.Limit == 0 || idx < params.Limit) {
			ret = append(ret, conn.ToStatus())
		}
		idx++
	}
	return ret
}

// Returns a Status representing the Connection with the provided ID
func (s *Service) GetByID(id uuid.UUID) *Status {
	if id == systemID {
		return systemStatus
	}
	if id == WASMID {
		return wasmStatus
	}
	conn, ok := s.connections[id]
	if !ok {
		s.Logger.Error(fmt.Sprintf("error getting connection by id [%v]", id))
		return nil
	}
	return conn.ToStatus()
}

// Total number of all connections managed by this service
func (s *Service) Count() int {
	return len(s.connections)
}


// Used by userless WASM messages
func (s *Service) SetWASMCallback(f func(string)) {
	s.wasmCallback = f
	err := s.OnOpen(WASMID)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("error processing WASM open event: %+v", err))
		return
	}
}

// Callback for when the backing connection is re-established
func (s *Service) OnOpen(connID uuid.UUID) error {
	if connID == WASMID {
		return s.onOpen(s, wasmConnection)
	}
	c, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	return s.onOpen(s, c)
}

// Sends a message to a provided Connection ID
func OnMessage(s *Service, connID uuid.UUID, message *Message) error {
	if connID == systemID {
		s.Logger.Warn("--- admin message received ---")
		s.Logger.Warn(fmt.Sprint(message))
		return nil
	}
	if connID == WASMID {
		return s.handler(s, wasmConnection, message.Svc, message.Cmd, message.Param)
	}
	c, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}

	return s.handler(s, c, message.Svc, message.Cmd, message.Param)
}


// Callback for when the backing connection is closed
func (s *Service) OnClose(connID uuid.UUID) error {
	c, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	return s.onOpen(s, c)
}
