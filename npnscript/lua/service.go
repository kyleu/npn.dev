package lua

import (
	lua "github.com/yuin/gopher-lua"
	"logur.dev/logur"
)

type Service struct {
	l      *lua.LState
	logger logur.Logger
}

func NewService(logger logur.Logger) *Service {
	l := lua.NewState()
	return &Service{l: l, logger: logger}
}

func (s *Service) Set(k string, v interface{}) {
	// TODO s.l.SetGlobal(k, "")
}

func (s *Service) Call(code string) (interface{}, error) {
	defer s.l.Close()
	err := s.l.DoString(code)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
