package project

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type Service struct {
	files  *npncore.FileLoader
	data   map[string]*Project
	logger logur.Logger
}

func NewService(files *npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{files: files, data: make(map[string]*Project), logger: logger}
}

func (s *Service) List() []string {
	return s.files.ListJSON("project")
}

func (s *Service) Summary(key string) (*Summary, error) {
	content, err := s.files.ReadFile("project/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find project file with key ["+key+"]")
	}
	tgt := &Summary{}
	err = npncore.FromJSON([]byte(content), tgt)
	if err != nil {
		return nil, err
	}
	return tgt, nil
}

func (s *Service) Load(key string) (*Project, error) {
	content, err := s.files.ReadFile("project/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find project file with key ["+key+"]")
	}
	tgt := &Project{}
	err = npncore.FromJSON([]byte(content), tgt)
	if err != nil {
		return nil, err
	}
	return tgt, nil
}

func (s *Service) Save(originalKey string, p *Project, overwrite bool) error {
	if len(originalKey) > 0 && originalKey != "new" {
		newProj, _ := s.Load(p.Key)
		if originalKey != p.Key && newProj != nil {
			return errors.New("remove the existing [" + originalKey + "] project before you overwrite it with this one")
		}
	}
	err := s.files.WriteFile("project/"+p.Key+".json", npncore.ToJSON(p, s.logger), overwrite)
	if err != nil {
		return errors.Wrap(err, "unable to write project")
	}
	if len(originalKey) > 0 && originalKey != "new" && originalKey != p.Key {
		err = s.Remove(originalKey)
		if err != nil {
			return errors.Wrap(err, "cannot remove original project")
		}
	}
	return nil
}

func (s *Service) Summaries() (Summaries, error) {
	projectKeys := s.List()
	summaries := make(Summaries, 0, len(projectKeys))
	for _, key := range projectKeys {
		sch, err := s.Summary(key)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, sch)
	}
	return summaries, nil
}

func (s *Service) Remove(key string) error {
	return errors.Wrap(s.files.Remove("project/"+key+".json"), "unable to remove project ["+key+"] file")
}
