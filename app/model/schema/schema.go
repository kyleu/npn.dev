package schema

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/util"
)

type Schema struct {
	Key      string    `json:"key"`
	Title    string    `json:"title"`
	Paths    Paths     `json:"paths"`
	Options  Options   `json:"options,omitempty"`
	Scalars  Scalars   `json:"scalars,omitempty"`
	Models   Models    `json:"models,omitempty"`
	Errors   []string  `json:"errors,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

func NewSchema(title string, paths []string, md *Metadata) *Schema {
	return &Schema{Key: util.Slugify(title), Title: title, Paths: paths, Metadata: md}
}

func (s *Schema) AddPath(path string) bool {
	if s.Paths.Exists(path) {
		return false
	}
	s.Paths = append(s.Paths, path)
	return true
}

func (s *Schema) AddOption(opt *Option) error {
	if s.Options.Get(opt.T, opt.K) != nil {
		return errors.New("option [" + opt.T + ":" + opt.K + "] already exists")
	}
	s.Options = append(s.Options, opt)
	return nil
}

func (s *Schema) AddScalar(sc *Scalar) error {
	if s.Scalars.Get(sc.Pkg, sc.Key) != nil {
		return errors.New("scalar [" + sc.Key + "] already exists")
	}
	s.Scalars = append(s.Scalars, sc)
	return nil
}

func (s *Schema) AddModel(m *Model) error {
	if s.Models.Get(m.Pkg, m.Key) != nil {
		return errors.New("model [" + m.Key + "] already exists")
	}
	s.Models = append(s.Models, m)
	return nil
}

func (s *Schema) Validate() *ValidationResult {
	return validateSchema(s)
}

func (s *Schema) ValidateModel(model *Model) *ValidationResult {
	r := &ValidationResult{Schema: s.Key}
	return validateModel(r, s, model)
}
