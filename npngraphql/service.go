package npngraphql

import (
	"context"
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"emperror.dev/errors"
	"logur.dev/logur"

	"github.com/graphql-go/graphql"
)

type Service struct {
	Logger logur.Logger
	cfg    graphql.SchemaConfig
	schema graphql.Schema
	app    npnweb.AppInfo
}

var graphQLService *Service

func NewService(app npnweb.AppInfo, queryName string, queries graphql.Fields, mutationName string, mutations graphql.Fields) (*Service, error) {
	logger := logur.WithFields(app.Logger(), map[string]interface{}{npncore.KeyService: npncore.KeyGraphQL})

	gqlQueryName = queryName
	gqlMutationName = mutationName

	schemaConfig := graphql.SchemaConfig{
		Query:        graphql.NewObject(graphql.ObjectConfig{Name: queryName, Fields: queries}),
		Mutation:     graphql.NewObject(graphql.ObjectConfig{Name: mutationName, Fields: mutations}),
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new schema")
	}
	svc := Service{Logger: logger, cfg: schemaConfig, schema: schema, app: app}

	logger.Debug("initialized GraphQL service")
	graphQLService = &svc
	return &svc, nil
}

func (s *Service) Run(operationName string, doc string, variables map[string]interface{}, ctx *npnweb.RequestContext) (*graphql.Result, error) {
	params := graphql.Params{
		Schema:         s.schema,
		RequestString:  doc,
		VariableValues: variables,
		OperationName:  operationName,
		Context:        context.WithValue(context.Background(), npncore.ContextKey, ctx),
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		errString := ""
		for i, e := range r.Errors {
			errString += fmt.Sprintf("%v: %v@%v", i, e.Message, e.Path)
			if i < len(r.Errors)-1 {
				errString += ", "
			}
		}
		return nil, errors.New("graphql errors [" + errString + "]")
	}
	return r, nil
}

func CtxF(f func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error)) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		c, ok := p.Context.Value(npncore.ContextKey).(*npnweb.RequestContext)
		if !ok {
			return nil, errors.New("no ctx in GraphQL resolve params")
		}
		return f(p, c)
	}
}
