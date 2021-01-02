package search

import (
	"github.com/gofrs/uuid"
	"logur.dev/logur"
)

type Service struct {
	logger logur.Logger
}

func (s Service) Run(p *Params, userID uuid.UUID, role string) (Results, error) {
	ret := Results{{Msg: "TODO"}}
	return ret, nil
}

func NewService(logger logur.Logger) *Service {
	return &Service{logger: logger}
}
