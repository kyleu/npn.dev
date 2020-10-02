package npnconnection

import (
	"encoding/json"
	"fmt"
	"sync"

	"emperror.dev/errors"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type Connection struct {
	ID      uuid.UUID
	Profile *npnuser.Profile
	Svc     string
	ModelID *uuid.UUID
	Channel *Channel
	socket  *websocket.Conn
	mu      sync.Mutex
}

func NewConnection(svc string, profile *npnuser.Profile, socket *websocket.Conn) *Connection {
	return &Connection{
		ID:      npncore.UUID(),
		Profile: profile,
		Svc:     svc,
		ModelID: nil,
		Channel: nil,
		socket:  socket,
	}
}

func (c *Connection) ToStatus() *Status {
	if c.Channel == nil {
		return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, ChannelSvc: npncore.KeySystem, ChannelID: nil}
	}
	return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, ChannelSvc: c.Channel.Svc, ChannelID: &c.Channel.ID}
}

func (c *Connection) Write(b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.socket.WriteMessage(websocket.TextMessage, b)
	return errors.Wrap(err, "unable to write to websocket")
}

func (c *Connection) Read() ([]byte, error) {
	_, message, err := c.socket.ReadMessage()
	return message, errors.Wrap(err, "unable to write to websocket")
}

func (c *Connection) Close() error {
	return c.socket.Close()
}

type Status struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"userID"`
	Username   string
	ChannelSvc string
	ChannelID  *uuid.UUID
}

type Statuses = []*Status

type Message struct {
	Svc   string          `json:"svc"`
	Cmd   string          `json:"cmd"`
	Param json.RawMessage `json:"param"`
}

func NewMessage(svc string, cmd string, param interface{}) *Message {
	return &Message{Svc: svc, Cmd: cmd, Param: json.RawMessage(npncore.ToJSON(param, nil))}
}

func (m *Message) String() string {
	return fmt.Sprintf("%s/%s", m.Svc, m.Cmd)
}

type OnlineUpdate struct {
	UserID    uuid.UUID `json:"userID"`
	Connected bool      `json:"connected"`
}

func DifferentPointerValues(l *uuid.UUID, r *uuid.UUID) bool {
	switch {
	case l != nil && r != nil:
		return *l != *r
	case l == nil && r != nil:
		return true
	case l != nil && r == nil:
		return true
	default:
		return false
	}
}
