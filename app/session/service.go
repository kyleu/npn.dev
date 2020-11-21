package session

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	files  npncore.FileLoader
	logger logur.Logger
}

func NewService(f npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{
		files:  f,
		logger: logger,
	}
}

func (s Service) Get(id uuid.UUID, sess string) (*Session, error) {
	ret := &Session{Key: sess}
	return ret, nil
}
