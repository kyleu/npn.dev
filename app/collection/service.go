package collection

import (
	"github.com/kyleu/npn/npnuser"
	"path"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	files  npncore.FileLoader
	logger logur.Logger
}

func NewService(f npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{files: f, logger: logger}
}

func (s *Service) List(userID *uuid.UUID) (Collections, error) {
	dirs := s.files.ListDirectories(dirFor(userID))

	ret := make(Collections, 0, len(dirs))
	for _, dir := range dirs {
		c, err := s.Load(userID, dir)
		if err != nil {
			return nil, err
		}
		if c == nil {
			c = &Collection{Key: dir}
		}
		ret = append(ret, c)
	}
	return ret, nil
}

func (s *Service) Counts(userID *uuid.UUID) (Summaries, error) {
	l, err := s.List(userID)
	if err != nil {
		return nil, err
	}

	ret := make(Summaries, 0, len(l))
	for _, coll := range l {
		count := len(s.files.ListJSON(path.Join(dirFor(userID), coll.Key, "requests")))
		ret = append(ret, &Summary{Coll: coll, Count: count})
	}
	return ret, nil
}

func (s *Service) Load(userID *uuid.UUID, key string) (*Collection, error) {
	p := path.Join(dirFor(userID), key)
	_, isDir := s.files.Exists(p)
	if !isDir {
		return nil, nil
	}
	ret := &Collection{}
	filePath := path.Join(p, "collection.json")
	fileExists, _ := s.files.Exists(filePath)
	if fileExists {
		content, err := s.files.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		err = npncore.FromJSON(content, ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read collection from ["+filePath+"]")
		}
	}

	return ret.Normalize(key, p), nil
}

func (s *Service) Delete(userID *uuid.UUID, key string) error {
	p := path.Join(dirFor(userID), key)
	return s.files.RemoveRecursive(p)
}

func dirFor(userID *uuid.UUID) string {
	if userID == nil || *userID == npnuser.SystemUserID {
		return "collections"
	}
	return path.Join("users", userID.String(), "collections")
}
