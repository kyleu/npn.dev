package schema

import (
	"encoding/json"
	"reflect"

	"emperror.dev/errors"
	"github.com/iancoleman/strcase"
	"github.com/kyleu/npn/app/util"
)

type ModelType struct {
	Key    string
	Title  string
	Plural string
}

var ModelTypeEnum = ModelType{Key: "enum", Title: "Enum", Plural: "Enums"}
var ModelTypeInput = ModelType{Key: "input", Title: "Input", Plural: "Inputs"}
var ModelTypeStruct = ModelType{Key: "struct", Title: "Struct", Plural: "Structs"}
var ModelTypeInterface = ModelType{Key: "interface", Title: "Interface", Plural: "Interfaces"}
var ModelTypeService = ModelType{Key: "service", Title: "Service", Plural: "Services"}
var ModelTypeUnion = ModelType{Key: "union", Title: "Union", Plural: "Unions"}
var ModelTypeIntersection = ModelType{Key: "intersection", Title: "Intersection", Plural: "Intersections"}

var AllModelTypes = []ModelType{
	ModelTypeEnum, ModelTypeInput, ModelTypeStruct,
	ModelTypeInterface, ModelTypeService, ModelTypeUnion, ModelTypeIntersection,
}

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
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = modelTypeFromString(s)
	return nil
}

type Model struct {
	Key         string    `json:"key"`
	Pkg         util.Pkg  `json:"pkg,omitempty"`
	Type        ModelType `json:"type"`
	Interfaces  []string  `json:"interfaces,omitempty"`
	Fields      Fields    `json:"fields,omitempty"`
	Indexes     Indexes   `json:"indexes,omitempty"`
	Description string    `json:"description,omitempty"`
	Metadata    *Metadata `json:"metadata,omitempty"`
}

func (m *Model) String() string {
	if len(m.Pkg) == 0 {
		return m.Key
	}
	return m.Pkg.StringWith(m.Key)
}

func (m *Model) ClassName(nr *util.NameRegistry) string {
	return nr.Replace(strcase.ToCamel(m.Key))
}

func (m *Model) PropName(nr *util.NameRegistry) string {
	return nr.Replace(strcase.ToLowerCamel(m.Key))
}

func (m *Model) AddField(f *Field) error {
	if f == nil {
		return errors.New("nil field")
	}
	if m.Fields.Get(f.Key) != nil {
		return errors.New(alreadyExists("field", f.Key))
	}
	m.Fields = append(m.Fields, f)
	return nil
}

func (m *Model) AddIndex(i *Index) error {
	if i == nil {
		return errors.New("nil index")
	}
	if m.Fields.Get(i.Key) != nil {
		return errors.New(alreadyExists("index", i.Key))
	}
	m.Indexes = append(m.Indexes, i)
	return nil
}

type Models []*Model

func (m Models) Get(pkg util.Pkg, key string) *Model {
	for _, x := range m {
		if reflect.DeepEqual(x.Pkg, pkg) && x.Key == key {
			return x
		}
	}
	return nil
}

func (m Models) HasField() bool {
	for _, model := range m {
		if len(model.Fields) > 0 {
			return true
		}
	}
	return false
}

func (m Models) HasIndex() bool {
	for _, model := range m {
		if len(model.Indexes) > 0 {
			return true
		}
	}
	return false
}

func (m Models) ByType(t ModelType) Models {
	ret := Models{}
	for _, x := range m {
		if x.Type == t {
			ret = append(ret, x)
		}
	}
	return ret
}
