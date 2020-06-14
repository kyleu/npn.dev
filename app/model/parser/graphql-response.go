package parser

import (
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
)

type GraphQLResponse struct {
	RootFile string         `json:"root"`
	Data     []interface{}  `json:"data"`
	Schema   *schema.Schema `json:"schema"`
}

func NewGraphQLResponse(key string) *GraphQLResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginGraphQL, Source: util.FilenameOf(key)}
	return &GraphQLResponse{
		RootFile: key,
		Data:     make([]interface{}, 0),
		Schema:   schema.NewSchema(key, []string{key}, &md),
	}
}
