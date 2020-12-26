package npnconnection

import (
	"encoding/json"
	"fmt"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

// Write a text message to the Connection matching the provided ID
func (s *Service) Write(connID uuid.UUID, message string) error {
	if connID == systemID {
		s.Logger.Warn("--- admin message sent ---")
		s.Logger.Warn(fmt.Sprint(message))
		return nil
	}
	if connID == WASMID {
		if s.wasmCallback == nil {
			return errors.New("no callback registered for WASM client")
		}
		s.wasmCallback(message)
		return nil
	}

	c, ok := s.connections[connID]
	if !ok {
		return errors.New("cannot load connection [" + connID.String() + "]")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.Wrap(err, "unable to write to websocket")
	}
	return nil
}

// Write a Message to the Connection matching the provided ID
func (s *Service) WriteMessage(connID uuid.UUID, message *Message) error {
	return s.Write(connID, npncore.ToJSON(message, s.Logger))
}

// Write a log message to the Connection matching the provided ID
func (s *Service) WriteLog(connID uuid.UUID, level string, msg string, ctx ...string) error {
	return s.WriteMessage(connID, NewMessage(npncore.KeySystem, npncore.KeyLog, NewLogMessage(level, msg, ctx...)))
}

// Broadcast a Message to call Connection instances
func (s *Service) Broadcast(message *Message, except ...uuid.UUID) error {
	// s.Logger.Debug(fmt.Sprintf("broadcasting message [%v::%v] to [%v] connections", message.Svc, message.Cmd, len(s.connections)))
	for id, _ := range s.connections {
		if !contains(except, id) {
			go func() {
				_ = s.Write(id, npncore.ToJSON(message, s.Logger))
			}()
		}
	}
	return nil
}

// Write a log message to all Connection instances
func (s *Service) BroadcastLog(level string, msg string, ctx ...string) error {
	return s.Broadcast(NewMessage(npncore.KeySystem, npncore.KeyLog, NewLogMessage(level, msg, ctx...)))
}

// Write a Message to the provided Channel
func (s *Service) WriteChannel(channel Channel, message *Message, except ...uuid.UUID) error {
	conns, ok := s.channels[channel]
	if !ok {
		return nil
	}

	// s.Logger.Debug(fmt.Sprintf("sending message [%v::%v] to [%v] connections", message.Svc, message.Cmd, len(conns)))
	for _, conn := range conns {
		if !contains(except, conn) {
			connID := conn

			go func() {
				_ = s.Write(connID, npncore.ToJSON(message, s.Logger))
			}()
		}
	}
	return nil
}

// Enter an loop that reads Message objects from the Connection matching the provided ID
func (s *Service) ReadLoop(connID uuid.UUID) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.New("cannot load connection [" + connID.String() + "]")
	}

	defer func() {
		_ = c.socket.Close()
		_, _ = s.Disconnect(connID)
		// s.Logger.Debug(fmt.Sprintf("closed websocket [%v]", connID.String()))
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			break
		}

		m := &Message{}
		err = json.Unmarshal(message, m)
		if err != nil {
			return errors.Wrap(err, "error decoding websocket message")
		}

		err = OnMessage(s, connID, m)
		if err != nil {
			_ = s.WriteLog(c.ID, "error", err.Error())
			s.Logger.Debug(fmt.Sprintf("error handling websocket message: %+v", err))
			// return errors.Wrap(err, "error handling websocket message")
		}
	}
	return nil
}
