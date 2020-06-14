package parser

import (
	"fmt"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"path"
)

type GraphQLParser struct {
	Key     string
	logger  logur.Logger
}

func (p *GraphQLParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := util.GetMatchingFiles(path.Join(root, "data", "graphql"), "*.graphql")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Origin: schema.OriginGraphQL})
	}
	return ret, nil
}

func (p *GraphQLParser) log(rsp *GraphQLResponse, err error) {
	if err != nil {
		rsp.Schema.Errors = append(rsp.Schema.Errors, err.Error())
		p.logger.Error(fmt.Sprintf("unable to parse GraphQL: %+v", err))
	}
}
