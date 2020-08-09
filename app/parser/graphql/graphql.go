package parsegraphql

import (
	"path"
	"strings"

	"emperror.dev/errors"

	parseutil "github.com/kyleu/npn/app/parser/util"
	"github.com/kyleu/npn/app/schema"
	"logur.dev/logur"
)

var gqlArgs = &struct{ IncludeDeprecated bool }{IncludeDeprecated: false}

type GraphQLParser struct {
	logger logur.Logger
}

func NewParser(logger logur.Logger) *GraphQLParser {
	logger = logur.WithFields(logger, map[string]interface{}{"service": schema.OriginGraphQL.Key})
	return &GraphQLParser{logger: logger}
}

func (p *GraphQLParser) Type() schema.Origin {
	return schema.OriginGraphQL
}

func (p *GraphQLParser) Detect(root string) ([]schema.DataSource, error) {
	gql, err := parseutil.GetMatchingFiles(path.Join(root, "data", "graphql"), "*.graphql")
	if err != nil {
		return nil, err
	}

	json, err := parseutil.GetMatchingFiles(path.Join(root, "data", "graphql"), "*.json")
	if err != nil {
		return nil, err
	}

	ret := make([]schema.DataSource, 0, len(gql)+len(json))

	for _, f := range gql {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Type: schema.OriginGraphQL})
	}

	for _, f := range json {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Type: schema.OriginGraphQL})
	}

	return ret, nil
}

func (p *GraphQLParser) IsValid(firstChars string) error {
	hasKeyword := false
	for _, x := range []string{"scalar", "enum", "type", "input"} {
		if strings.Contains(firstChars, x) {
			hasKeyword = true
			break
		}
	}
	if !hasKeyword {
		return errors.New("not GraphQL")
	}
	return nil
}
