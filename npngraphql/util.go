package npngraphql

import (
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"logur.dev/logur"
)

type Callback func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error)

func ParamSetFromGraphQLParams(key string, params graphql.ResolveParams, logger logur.Logger) *npncore.Params {
	orderings := make(npncore.Orderings, 0)
	o, ok := params.Args["orders"]
	if ok {
		for _, x := range o.([]interface{}) {
			m := x.(map[string]interface{})
			col := npncore.MapGetString(m, "col", logger)
			asc := npncore.MapGetBool(m, "asc", logger)
			var defaultOrdering = npncore.Orderings{{Column: col, Asc: asc}}
			orderings = append(orderings, defaultOrdering...)
		}
	}

	limit := 0
	l, ok := params.Args["limit"]
	if ok {
		limit = l.(int)
	}

	offset := 0
	x, ok := params.Args["offset"]
	if ok {
		offset = x.(int)
	}

	ret := &npncore.Params{Key: key, Orderings: orderings, Limit: limit, Offset: offset}
	return ret.Filtered(logger)
}

func ErrorResponseJSON(logger logur.Logger, errors ...error) *graphql.Result {
	var errs = make([]gqlerrors.FormattedError, 0, len(errors))

	for _, err := range errors {
		logger.Warn(fmt.Sprintf("error running GraphQL: %+v", err))
		errs = append(errs, gqlerrors.FormattedError{Message: err.Error()})
	}

	return &graphql.Result{
		Errors: errs,
	}
}

var OrderingInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Ordering",
		Fields: graphql.InputObjectConfigFieldMap{
			"col": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"asc": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
	},
)

var ListArgs = graphql.FieldConfigArgument{
	"orders": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(OrderingInputType)),
	},
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var IDArgs = graphql.FieldConfigArgument{
	npncore.KeyID: &graphql.ArgumentConfig{
		Type: ScalarUUID,
	},
}

var KeyArgs = graphql.FieldConfigArgument{
	npncore.KeyKey: &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
