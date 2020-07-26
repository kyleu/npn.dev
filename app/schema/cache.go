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

func NewService(logger logur.Logger) *Service {
	return &Service{files: npncore.NewFileLoader(logger), data: make(map[string]*Schema), logger: logger}
}

func (s *Service) List() []string {
	return s.files.ListJSON("schema")
}

func (s *Service) Summary(key string) (*Summary, error) {
	content, err := s.files.ReadFile("schema/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key ["+key+"]")
	}
	tgt := &Summary{}
	npncore.FromJSON([]byte(content), tgt, s.logger)
	return tgt, nil
}

func (s *Service) Load(key string) (*Schema, error) {
	content, err := s.files.ReadFile("schema/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key ["+key+"]")
	}
	tgt := &Schema{}
	npncore.FromJSON([]byte(content), tgt, s.logger)
	return tgt, nil
}

func (s *Service) Save(sch *Schema, overwrite bool) error {
	return s.files.WriteFile("schema/"+sch.Key+".json", npncore.ToJSON(sch, s.logger), overwrite)
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
