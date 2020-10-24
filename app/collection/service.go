package collection

import (
	"os"
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

func (s *Service) Save(userID *uuid.UUID, originalKey string, newKey string, title string, description string) error {
	originalKey = npncore.Slugify(originalKey)
	newKey = npncore.Slugify(newKey)

	var orig *Collection
	var err error

	if len(originalKey) > 0 {
		orig, err = s.Load(userID, originalKey)
		if err != nil {
			return errors.Wrap(err, "unable to load original collection ["+originalKey+"]")
		}
		if orig != nil && originalKey != newKey {
			o := path.Join(s.files.Root(), dirFor(userID), originalKey)
			n := path.Join(s.files.Root(), dirFor(userID), newKey)
			err := os.Rename(o, n)
			if err != nil {
				return errors.Wrap(err, "unable to rename original collection ["+originalKey+"] in path ["+o+"]")
			}
		}
	}

	n := &Collection{
		Key:         newKey,
		Title:       title,
		Description: description,
	}

	if orig == nil {
		n.Owner = "system"
	} else {
		n.Owner = orig.Owner
		n.RequestOrder = orig.RequestOrder
	}
	n.Path = newKey

	p := path.Join(dirFor(userID), newKey, "collection.json")
	content := npncore.ToJSON(n, s.logger)
	err = s.files.WriteFile(p, []byte(content), true)
	if err != nil {
		return errors.Wrap(err, "unable to save collection ["+newKey+"]")
	}

	return nil
}

func (s *Service) Delete(userID *uuid.UUID, key string) error {
	p := path.Join(dirFor(userID), key)
	return s.files.RemoveRecursive(p)
}

var systemUserID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000")

func dirFor(userID *uuid.UUID) string {
	if userID == nil || *userID == systemUserID {
		return "collections"
	}
	return path.Join("users", userID.String(), "collections")
}
