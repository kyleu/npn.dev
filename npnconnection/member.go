package npnconnection

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
)

func (s *Service) GetOnline(ch Channel) []uuid.UUID {
	connections, ok := s.channels[ch]
	if !ok {
		connections = make([]uuid.UUID, 0)
	}
	online := make([]uuid.UUID, 0)
	for _, cID := range connections {
		c, ok := s.connections[cID]
		if ok && c != nil && (!contains(online, c.Profile.UserID)) {
			online = append(online, c.Profile.UserID)
		}
	}

	return online
}

func (s *Service) sendOnlineUpdate(ch Channel, connID uuid.UUID, userID uuid.UUID, connected bool) error {
	p := OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := NewMessage(npncore.KeySystem, "online-update", p)
	return s.WriteChannel(ch, onlineMsg, connID)
}
