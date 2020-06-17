package parsejsonschema

import (
	"github.com/kyleu/npn/app/model/schema"
	"github.com/santhosh-tekuri/jsonschema/v2"
)

type JSONSchemaResponse struct {
	RootFile string         `json:"root"`
	Data     []interface{}  `json:"data"`
	Schema   *schema.Schema `json:"schema"`
}

func NewJSONSchemaResponse(paths []string) *JSONSchemaResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginJSONSchema, Source: paths[0]}
	return &JSONSchemaResponse{
		RootFile: paths[0],
		Data:     make([]interface{}, 0),
		Schema:   schema.NewSchema(paths[0], paths, &md),
	}
}

type JSONSchema struct {
	js     *jsonschema.Schema `json:"root"`
}

func (j *JSONSchema) debug() string {
	return "Hello!"
}
