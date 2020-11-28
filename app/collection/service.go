package collection

import (
	"path"

	"github.com/kyleu/npn/npnuser"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	multiuser bool
	files     npncore.FileLoader
	logger    logur.Logger
}

func NewService(multiuser bool, f npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{multiuser: multiuser, files: f, logger: logger}
}

func (s *Service) List(userID *uuid.UUID) (Collections, error) {
	d := s.dirFor(userID)
	dExists, isDir := s.files.Exists(d)
	if (!dExists) || (!isDir) {
		return make(Collections, 0), nil
	}

	dirs := s.files.ListDirectories(d)

	if len(dirs) == 0 {
		return Collections{defaultCollection}, nil
	}

	ret := make(Collections, 0, len(dirs))
	hasDefault := false
	for _, dir := range dirs {
		c, err := s.Load(userID, dir)
		if err != nil {
			return nil, err
		}
		if c == nil {
			c = &Collection{Key: dir}
		}
		if dir == "_" {
			hasDefault = true
		}
		ret = append(ret, c)
	}
	if !hasDefault {
		ret = append(Collections{defaultCollection}, ret...)
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
		count := len(s.files.ListJSON(path.Join(s.dirFor(userID), coll.Key, "requests")))
		ret = append(ret, &Summary{Coll: coll, Count: count})
	}
	return ret, nil
}

func (s *Service) Load(userID *uuid.UUID, key string) (*Collection, error) {
	p := path.Join(s.dirFor(userID), key)
	_, isDir := s.files.Exists(p)
	if !isDir {
		if key == "_" {
			return defaultCollection.Normalize(key, p), nil
		}
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
	} else if key == "_" {
		return defaultCollection.Normalize(key, p), nil
	}

	return ret.Normalize(key, p), nil
}

func (s *Service) Delete(userID *uuid.UUID, key string) error {
	p := path.Join(s.dirFor(userID), key)
	return s.files.RemoveRecursive(p)
}

func (s *Service) dirFor(userID *uuid.UUID) string {
	if (!s.multiuser) || userID == nil || *userID == npnuser.SystemUserID {
		return "collections"
	}
	return path.Join("users", userID.String(), "collections")
}
