package parseutil

import (
	"github.com/kyleu/npn/app/schema"
)

type ParseResponse struct {
	Data   []interface{}  `json:"data"`
	Schema *schema.Schema `json:"schema,omitempty"`
}

func NewParseResponse(paths []string, md schema.Metadata) *ParseResponse {
	return &ParseResponse{
		Data:   make([]interface{}, 0),
		Schema: schema.NewSchema(paths[0], paths, &md),
	}
}
