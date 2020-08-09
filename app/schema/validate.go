package schema

import (
	"github.com/kyleu/npn/app/schema/schematypes"
)

const (
	LevelInfo = iota
	LevelWarn
	LevelError
)

type ValidationMessage struct {
	Category string
	ModelKey string
	Msg      string
	Level    int
}

type ValidationResult struct {
	Schema   string
	Messages []ValidationMessage
	Duration int64
}

func (v *ValidationResult) log(category string, modelKey string, msg string, level int) {
	v.Messages = append(v.Messages, ValidationMessage{Category: category, ModelKey: modelKey, Msg: msg, Level: level})
}

func validateSchema(s *Schema) *ValidationResult {
	r := &ValidationResult{Schema: s.Key}
	for _, m := range s.Models {
		r = validateModel(r, s, m)
	}
	return r
}

func validateModel(r *ValidationResult, s *Schema, m *Model) *ValidationResult {
	encountered := map[string]bool{}
	for _, f := range m.Fields {
		if encountered[f.Key] {
			r.log(m.Type.Key, m.Key, m.Type.String()+" ["+m.Key+"] field ["+f.Key+"] appears twice", LevelError)
		}
		encountered[f.Key] = true
	}
	for _, v := range m.Fields {
		validateType(r, s, "model", m.Key, v.Key, v.Type)
	}
	return r
}

func validateType(r *ValidationResult, s *Schema, mType string, mKey string, fKey string, f schematypes.Type) {
	switch t := f.(type) {
	case schematypes.Wrapped:
		validateType(r, s, mType, mKey, fKey, t.V)
	case schematypes.Unknown:
		r.log(mType, mKey, "field ["+fKey+"] has unknown type ["+t.X+"]", LevelWarn)
	case schematypes.Error:
		r.log(mType, mKey, "field ["+fKey+"] has error: "+t.Message, LevelWarn)
	case schematypes.List:
		validateType(r, s, mType, mKey, fKey, t.T)
	case schematypes.Option:
		validateType(r, s, mType, mKey, fKey, t.T)
	case schematypes.Reference:
		if s.Models.Get(t.Pkg, t.T) == nil && s.Scalars.Get(t.Pkg, t.T) == nil {
			r.log(mType, mKey, "field ["+fKey+"] has reference to unknown type ["+t.Pkg.String()+"::"+t.T+"]", LevelWarn)
		}
	}
}
