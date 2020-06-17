package parsegraphql

import (
	"github.com/kyleu/npn/app/model/schema"
)

type GraphQLResponse struct {
	RootFile string         `json:"root"`
	Data     []interface{}  `json:"data"`
	Schema   *schema.Schema `json:"schema"`
}

func NewGraphQLResponse(paths []string) *GraphQLResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginGraphQL, Source: paths[0]}
	return &GraphQLResponse{
		RootFile: paths[0],
		Data:     make([]interface{}, 0),
		Schema:   schema.NewSchema(paths[0], paths, &md),
	}
}
