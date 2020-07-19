package parsegraphql

import (
	"emperror.dev/errors"
	"github.com/graph-gophers/graphql-go/introspection"
	parseutil "github.com/kyleu/npn/app/parser/util"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/schema/schematypes"
	"logur.dev/logur"
)

func parseGraphQLType(ret *parseutil.ParseResponse, t *introspection.Type, logger logur.Logger) error {
	key := *t.Name()
	md := getGraphQLMetadata(t.Description())
	switch t.Kind() {
	case "SCALAR":
		return parseScalar(ret.Schema, key, md)
	case "ENUM":
		return parseEnum(ret.Schema, t, key, md)
	case "INTERFACE":
		return parseInterface(ret.Schema, t, key, md, logger)
	case "OBJECT":
		return parseObject(ret.Schema, t, key, md, logger)
	case "INPUT_OBJECT":
		return parseInputObject(ret.Schema, t, key, md, logger)
	case "UNION":
		return parseUnion(ret.Schema, t, key, md)
	default:
		return errors.New("invalid graphql kind [" + t.Kind() + "]")
	}
}

func getGraphQLMetadata(desc *string, comments ...string) *schema.Metadata {
	if desc == nil {
		d := ""
		desc = &d
	}
	return &schema.Metadata{
		Description: *desc,
		Comments:    comments,
		Origin:      schema.OriginGraphQL,
	}
}

func getGraphQLInterfaces(t *introspection.Type) []string {
	interfaces := t.Interfaces()
	if interfaces == nil {
		return nil
	}
	ret := make([]string, 0, len(*interfaces))
	for _, iface := range *interfaces {
		ret = append(ret, *iface.Name())
	}
	return ret
}

func getGraphQLFields(t *introspection.Type, logger logur.Logger) schema.Fields {
	gqlFields := t.Fields(gqlArgs)
	if gqlFields == nil {
		return nil
	}
	var fields schema.Fields
	for _, field := range *gqlFields {
		fields = append(fields, &schema.Field{Key: field.Name(), Type: getGraphQLType(field.Type(), field.Args(), logger)})
	}
	return fields
}

func getGraphQLInputFields(t *introspection.Type, logger logur.Logger) schema.Fields {
	gqlFields := t.InputFields()
	if gqlFields == nil {
		return nil
	}
	var fields schema.Fields
	for _, field := range *gqlFields {
		fields = append(fields, &schema.Field{Key: field.Name(), Type: getGraphQLType(field.Type(), nil, logger)})
	}
	return fields
}

func getGraphQLType(t *introspection.Type, args []*introspection.InputValue, logger logur.Logger) schematypes.Wrapped {
	if t == nil {
		return schematypes.NilWrapped
	}
	if t.Name() != nil {
		if len(args) > 0 {
			return funcCall(t, args, logger)
		}
		return getTypeForGraphQLName(*t.Name())
	}
	switch t.Kind() {
	case "NON_NULL":
		return unwrapGraphQLOption(getGraphQLType(t.OfType(), nil, logger))
	case "LIST":
		typ := schematypes.List{T: getGraphQLType(t.OfType(), nil, logger)}
		return schematypes.Wrap(typ)
	default:
		return schematypes.OptionWrapped(schematypes.Reference{T: t.Kind()})
	}
}

func unwrapGraphQLOption(ret schematypes.Type) schematypes.Wrapped {
	switch t := ret.(type) {
	case schematypes.Wrapped:
		return unwrapGraphQLOption(t.V)
	case schematypes.Option:
		return t.T
	default:
		return schematypes.Wrap(ret)
	}
}

func getTypeForGraphQLName(name string) schematypes.Wrapped {
	var ret schematypes.Type
	switch name {
	case "Boolean", "Bool":
		ret = schematypes.Bool{}
	case "Int":
		ret = schematypes.Int{}
	case "Float":
		ret = schematypes.Float{}
	case "String", "ID":
		ret = schematypes.String{}
	default:
		ret = schematypes.Reference{T: name}
	}
	return schematypes.Wrap(ret)
}

func funcCall(t *introspection.Type, args []*introspection.InputValue, logger logur.Logger) schematypes.Wrapped {
	schemaArgs := make(schematypes.Arguments, 0, len(args))
	for _, arg := range args {
		schemaArgs = append(schemaArgs, schematypes.Argument{
			Key:  arg.Name(),
			Type: getGraphQLType(arg.Type(), nil, logger),
		})
	}
	return schematypes.Wrap(schematypes.Method{Args: schemaArgs, Ret: getGraphQLType(t, nil, logger)})
}
