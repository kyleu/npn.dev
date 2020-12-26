package npnconnection

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// A group or room used to send broadcast messages
type Channel struct {
	Svc string
	ID  uuid.UUID
}

// Format "svc:id"
func (ch *Channel) String() string {
	return fmt.Sprintf("%s:%s", ch.Svc, ch.ID)
}

// Adds a Connection to this Channel
func (s *Service) Join(connID uuid.UUID, ch Channel) error {
	conn, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	if conn.Channel != &ch {
		conn.Channel = &ch
	}

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	if !contains(curr, connID) {
		s.channels[ch] = append(curr, connID)
	}
	return nil
}

// Removes a Connection from this Channel
func (s *Service) Leave(connID uuid.UUID, ch Channel) error {
	conn, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	conn.Channel = nil

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	filtered := make([]uuid.UUID, 0)
	for _, i := range curr {
		if i != connID {
			filtered = append(filtered, i)
		}
	}

	if len(filtered) == 0 {
		delete(s.channels, ch)
		return nil
	}

	s.channels[ch] = filtered
	return s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, false)
}

func contains(s []uuid.UUID, e uuid.UUID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
