package collection

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
	"path"
	"strings"
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
		err = npncore.FromJSON([]byte(content), ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read collection from ["+filePath+"]")
		}
	}

	return ret.Normalize(key, p), nil
}

func (s *Service) Requests(key string) ([]string, error) {
	p := path.Join(rootDir, key, "requests")
	return s.files.ListJSON(p), nil
}

func (s *Service) LoadRequest(c string, f string) (*request.Request, error) {
	f = strings.TrimSuffix(f, ".json")
	p := path.Join(rootDir, c, "requests", f + ".json")
	content, err := s.files.ReadFile(p)
	if err != nil {
		return nil, err
	}
	ret := &request.Request{}
	err = npncore.FromJSON([]byte(content), ret)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read request from ["+p+"]")
	}

	return ret.Normalize(f), nil
}
