package schema

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/kyleu/npn/app/util"
	"strings"
)

type ModelType struct {
	Key string
}

var ModelTypeEnum = ModelType{Key: "enum"}
var ModelTypeInput = ModelType{Key: "input"}
var ModelTypeStruct = ModelType{Key: "struct"}
var ModelTypeInterface = ModelType{Key: "interface"}
var ModelTypeService = ModelType{Key: "service"}
var ModelTypeUnion = ModelType{Key: "union"}

var AllModelTypes = []ModelType{ModelTypeEnum, ModelTypeInput, ModelTypeStruct, ModelTypeInterface, ModelTypeService, ModelTypeUnion}

func modelTypeFromString(s string) ModelType {
	for _, t := range AllModelTypes {
		if t.Key == s {
			return t
		}
	}
	return ModelTypeStruct
}

func (t *ModelType) String() string {
	return t.Key
}

func (t *ModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *ModelType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = modelTypeFromString(s)
	return nil
}

type Model struct {
	Key        string    `json:"key"`
	Pkg        []string  `json:"pkg"`
	Type       ModelType `json:"type"`
	Interfaces []string  `json:"interfaces,omitempty"`
	Fields     Fields    `json:"fields"`
	Metadata   *Metadata `json:"metadata,omitempty"`
}

func (m *Model) ID() string {
	if len(m.Pkg) == 0 {
		return m.Key
	}
	return strings.Join(append(m.Pkg, m.Key), ".")
}

func (m *Model) AddField(f *Field) error {
	if m.Fields.Get(f.Key) != nil {
		return errors.New("field [" + f.Key + "] already exists")
	}
	m.Fields = append(m.Fields, f)
	return nil
}

type Models []*Model

func (s Models) Get(pkg []string, key string) *Model {
	for _, x := range s {
		if util.StringArraysEqual(x.Pkg, pkg) && x.Key == key {
			return x
		}
	}
	return nil
}
