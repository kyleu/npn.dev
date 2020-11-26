package session

import (
	"path"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"logur.dev/logur"
)

type Service struct {
	multiuser bool
	files  npncore.FileLoader
	logger logur.Logger
}

func NewService(multiuser bool, f npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{
		multiuser: multiuser,
		files:  f,
		logger: logger,
	}
}

func (s *Service) List(userID *uuid.UUID) (Sessions, error) {
	jsons := s.files.ListJSON(s.dirFor(userID))

	if len(jsons) == 0 {
		return Sessions{{Key: "_", Title: "Default Session"}}, nil
	}

	ret := make(Sessions, 0, len(jsons))
	for _, json := range jsons {
		c, err := s.Load(userID, json)
		if err != nil {
			return nil, err
		}
		if c == nil {
			c = &Session{Key: json}
		}
		ret = append(ret, c)
	}
	return ret, nil
}

func (s *Service) Load(userID *uuid.UUID, key string) (*Session, error) {
	if len(key) == 0 {
		key = "_"
	}
	p := path.Join(s.dirFor(userID), key+".json")
	_, exists := s.files.Exists(p)
	if !exists {
		if key == "_" {
			return defaultSession, nil
		}
		return nil, nil
	}
	ret := &Session{}
	filePath := p
	fileExists, _ := s.files.Exists(filePath)
	if fileExists {
		content, err := s.files.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		err = npncore.FromJSON(content, ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read session from ["+filePath+"]")
		}
	}

	return ret.Normalize(key), nil
}

func (s *Service) Counts(userID *uuid.UUID) ([]*Summary, error) {
	l, err := s.List(userID)
	if err != nil {
		return nil, err
	}

	ret := make(Summaries, 0, len(l))
	for _, coll := range l {
		ret = append(ret, coll.ToSummary())
	}
	return ret, nil
}

func (s *Service) dirFor(userID *uuid.UUID) string {
	if (s.multiuser) || userID == nil || *userID == npnuser.SystemUserID {
		return "sessions"
	}
	return path.Join("users", userID.String(), "sessions")
}
