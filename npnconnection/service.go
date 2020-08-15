package npnconnection

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/kyleu/npn/npncore"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npnservice/user"

	"github.com/gofrs/uuid"
	"logur.dev/logur"
)

type Handler func(s *Service, conn *Connection, svc string, cmd string, param json.RawMessage) error

type Service struct {
	connections   map[uuid.UUID]*Connection
	connectionsMu sync.Mutex
	channels      map[Channel][]uuid.UUID
	channelsMu    sync.Mutex
	Logger        logur.Logger
	Users         *user.Service
	Auths         *auth.Service
	handler       Handler
	Context       interface{}
}

func NewService(logger logur.Logger, users *user.Service, auths *auth.Service, handler Handler, ctx interface{}) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeySocket})
	return &Service{
		connections:   make(map[uuid.UUID]*Connection),
		connectionsMu: sync.Mutex{},
		channels:      make(map[Channel][]uuid.UUID),
		channelsMu:    sync.Mutex{},
		Logger:        logger,
		Users:         users,
		Auths:         auths,
		handler:       handler,
		Context:       ctx,
	}
}

var systemID = uuid.FromStringOrNil("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF")
var systemStatus = Status{ID: systemID, UserID: systemID, Username: "System Broadcast", ChannelSvc: npncore.KeySystem, ChannelID: &systemID}

func (s *Service) List(params *npncore.Params) Statuses {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyConnection, params)
	ret := make(Statuses, 0)
	ret = append(ret, &systemStatus)
	var idx = 0
	for _, conn := range s.connections {
		if idx >= params.Offset && (params.Limit == 0 || idx < params.Limit) {
			ret = append(ret, conn.ToStatus())
		}
		idx++
	}
	return ret
}

func (s *Service) GetByID(id uuid.UUID) *Status {
	if id == systemID {
		return &systemStatus
	}
	conn, ok := s.connections[id]
	if !ok {
		s.Logger.Error(fmt.Sprintf("error getting connection by id [%v]", id))
		return nil
	}
	return conn.ToStatus()
}

func (s *Service) Count() int {
	return len(s.connections)
}

func onMessage(s *Service, connID uuid.UUID, message Message) error {
	if connID == systemID {
		s.Logger.Warn("--- admin message received ---")
		s.Logger.Warn(fmt.Sprint(message))
		return nil
	}
	c, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}

	return s.handler(s, c, message.Svc, message.Cmd, message.Param)
}
