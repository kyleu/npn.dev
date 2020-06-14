package parser

import (
	"emperror.dev/errors"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/kyleu/npn/app/model/schema"
	"io/ioutil"
)

func (p *GraphQLParser) ParseSchemaFile(path string) (*GraphQLResponse, error) {
	return p.parse(path, NewGraphQLResponse(path))
}

func (p *GraphQLParser) parse(fn string, ret *GraphQLResponse) (*GraphQLResponse, error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	schStr := string(b)
	sch, err := graphql.ParseSchema(schStr, nil)
	if err != nil {
		return nil, err
	}
	intro := sch.Inspect()
	for _, t := range intro.Types() {
		err = parseType(ret, t)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing type [%v] from file [" + fn + "]")
		}
	}
	for _, d := range intro.Directives() {
		md := getGraphQLMetadata(d.Description())
		err = ret.Schema.AddOption(&schema.Option{T: "directive", K: d.Name(), V: d.Name(), Metadata: md})
		if err != nil {
			return nil, errors.Wrap(err, "unable to add option")
		}
	}
	return ret, nil
}

func parseType(ret *GraphQLResponse, t *introspection.Type) error {
	ret.Data = append(ret.Data, map[string]interface{}{
		"name":     t.Name(),
		"kind":     t.Kind(),
		"ifaces":   t.Interfaces(),
		"possible": t.PossibleTypes(),
	})

	key := *t.Name()
	md := getGraphQLMetadata(t.Description())
	switch t.Kind() {
	case "SCALAR":
		return parseScalar(ret.Schema, t, key, md)
	case "ENUM":
		return parseEnum(ret.Schema, t, key, md)
	case "INTERFACE":
		return parseInterface(ret.Schema, t, key, md)
	case "OBJECT":
		return parseObject(ret.Schema, t, key, md)
	case "INPUT_OBJECT":
		return parseInputObject(ret.Schema, t, key, md)
	case "UNION":
		return parseUnion(ret.Schema, t, key, md)
	default:
		return errors.New("invalid graphql kind [" + t.Kind() + "]")
	}
}


func parseScalar(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	return s.AddScalar(&schema.Scalar{Key: key, Type: key, Metadata: md})
}

func parseEnum(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	var vals schema.EnumValues
	evs := t.EnumValues(nil)
	if evs != nil {
		for _, ev := range *evs {
			md := getGraphQLMetadata(ev.Description())
			vals = append(vals, &schema.EnumValue{Key: ev.Name(), Metadata: md})
		}
	}
	ret := &schema.Enum{Key: key, Values: vals, Metadata: md}
	return s.AddEnum(ret)
}

func parseInterface(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	ret := &schema.Model{Key: key, Fields: getGraphQLFields(t), Metadata: md}
	return s.AddModel(ret)
}

func parseObject(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	ret := &schema.Model{Key: key, Fields: getGraphQLFields(t), Metadata: md}
	return s.AddModel(ret)
}

func parseInputObject(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	ret := &schema.Model{Key: key, Fields: getGraphQLFields(t), Metadata: md}
	return s.AddModel(ret)
}

func parseUnion(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	ret := &schema.Union{Key: key, Metadata: md}
	return s.AddUnion(ret)
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
		// Source:   util.FilenameOf(pos.Filename),
		// Line:     pos.Line,
		// Column:   pos.Column - 1,
	}
}

func getGraphQLFields(t *introspection.Type) schema.Fields {
	var fields schema.Fields
	gqlFields := t.Fields(nil)
	if gqlFields != nil {
		for _, field := range *gqlFields {
			fields = append(fields, &schema.Field{Key: field.Name(), Type: getGraphQLType(field.Type())})
		}
	}
	return fields
}

func getGraphQLType(t *introspection.Type) string {
	if t == nil {
		return "--nil"
	}
	if t.Name() != nil {
		return *t.Name()
	}
	switch t.Kind() {
	case "NON_NULL":
		return getGraphQLType(t.OfType())
	case "LIST":
		return "[]" + getGraphQLType(t.OfType())
	default:
		return "*" + t.Kind()
	}
}
