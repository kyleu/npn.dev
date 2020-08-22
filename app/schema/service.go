package schema

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	files  *npncore.FileLoader
	data   map[string]*Schema
	logger logur.Logger
}

const schemaPath = "schema/"

func NewService(files *npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{files: files, data: make(map[string]*Schema), logger: logger}
}

func (s *Service) List() []string {
	return s.files.ListJSON("schema")
}

func (s *Service) Summary(key string) (*Summary, error) {
	content, err := s.files.ReadFile(schemaPath + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key ["+key+"]")
	}
	tgt := &Summary{}
	err = npncore.FromJSON([]byte(content), tgt)
	if err != nil {
		return nil, err
	}
	return tgt, nil
}

func (s *Service) Load(key string) (*Schema, error) {
	content, err := s.files.ReadFile(schemaPath + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key ["+key+"]")
	}
	tgt := &Schema{}
	err = npncore.FromJSON([]byte(content), tgt)
	if err != nil {
		return nil, err
	}
	return tgt, nil
}

func (s *Service) LoadAll(keys []string) (Schemata, error) {
	var schemata Schemata
	for _, schemaKey := range keys {
		sch, err := s.Load(schemaKey)
		if err != nil {
			err = errors.Wrap(err, "cannot load schema ["+schemaKey+"]")
			return nil, err
		}
		schemata = append(schemata, sch)
	}
	return schemata, nil
}

func (s *Service) Save(sch *Schema, overwrite bool) error {
	return s.files.WriteFile(schemaPath+sch.Key+".json", npncore.ToJSON(sch, s.logger), overwrite)
}

func (s *Service) Summaries() (Summaries, error) {
	schemaKeys := s.List()
	ret := make(Summaries, 0, len(schemaKeys))
	for _, key := range schemaKeys {
		sch, err := s.Summary(key)
		if err != nil {
			return nil, err
		}
		ret = append(ret, sch)
	}
	return ret, nil
}
