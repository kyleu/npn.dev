package imprt

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"io/ioutil"
	"logur.dev/logur"
	"mime/multipart"
	"path"
)

type Service struct {
	files  *npncore.FileLoader
	logger logur.Logger
}

func NewService(files *npncore.FileLoader, logger logur.Logger) *Service {
	return &Service{files: files, logger: logger}
}

func (s *Service) Create(key string, files []File) error {
	cfg := Config{Files: files, Status: "created"}
	json := npncore.ToJSON(cfg, s.logger)
	p := path.Join("import", key, "_import.json")
	err := s.files.WriteFile(p, []byte(json), false)
	if err != nil {
		return errors.Wrap(err, "cannot write import summary")
	}
	return nil
}

func (s *Service) Load(key string) (*Config, Outputs, error) {
	p := path.Join("import", key, "_import.json")
	content, err := s.files.ReadFile(p)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot read import summary")
	}
	cfg := &Config{}
	err = npncore.FromJSON(content, cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot parse import summary JSON")
	}

	outs := make(Outputs, 0, len(cfg.Files))
	for _, f := range cfg.Files {
		outs = append(outs, s.LoadFile(key, f.Filename, f.ContentType))
	}

	return cfg, outs, nil
}

func (s *Service) LoadFile(key string, filename string, contentType string) *Output {
	p := path.Join("import", key, "files", filename)
	content, err := s.files.ReadFile(p)
	if err != nil {
		return &Output{Filename: filename, Type: contentType, Value: len(content), Error: errors.Wrap(err, "cannot read import summary").Error()}
	}
	t, o, err := parse(filename, contentType, content)
	if err != nil {
		return &Output{Filename: filename, Type: t, Value: o, Error: errors.Wrap(err, "cannot parse ["+filename+"]").Error()}
	}

	return &Output{Filename: filename, Type: t, Value: o}
}

func (s *Service) WriteFile(key string, filename string, f multipart.File) error {
	p := path.Join("import", key, "files", filename)
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "unable to read file ["+p+"]")
	}
	return s.files.WriteFile(p, []byte(content), true)
}
