package collection

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
	"os"
	"path"
)

const rootDir = "collections"

type Service struct {
	files  *npncore.FileLoader
	logger logur.Logger
}

func NewService(f *npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{files: f, logger: logger}
}

func (s *Service) List() (Collections, error) {
	dirs := s.files.ListDirectories(rootDir)

	ret := make(Collections, 0, len(dirs))
	for _, dir := range dirs {
		c, err := s.Load(dir)
		if err != nil {
			return nil, err
		}
		ret = append(ret, c)
	}
	return ret, nil
}

func (s *Service) Load(key string) (*Collection, error) {
	p := path.Join(rootDir, key)
	_, isDir := s.files.Exists(p)
	if !isDir {
		return nil, errors.New("collection [" + key + "] does not exist")
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

func (s *Service) Save(originalKey string, newKey string, title string, description string) error {
	orig, err := s.Load(originalKey)

	if orig != nil && originalKey != newKey {
		o := path.Join(s.files.Root(), rootDir, originalKey)
		n := path.Join(s.files.Root(), rootDir, newKey)
		err := os.Rename(o, n)
		if err != nil {
			return errors.Wrap(err, "unable to rename original collection [" + originalKey + "] in path [" + o + "]")
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
	}
	n.Path = newKey

	p := path.Join(rootDir, newKey, "collection.json")
	content := npncore.ToJSON(n, s.logger)
	err = s.files.WriteFile(p, []byte(content), true)
	if err != nil {
		return errors.Wrap(err, "unable to save collection [" + newKey + "]")
	}

	return nil
}

func (s *Service) Delete(key string) error {
	p := path.Join(rootDir, key)
	return s.files.RemoveRecursive(p)
}
