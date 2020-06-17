package parsegraphql

import (
	"emperror.dev/errors"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"io/ioutil"
	"logur.dev/logur"
)

func (p *GraphQLParser) ParseSchemaFile(paths []string) (*GraphQLResponse, error) {
	return p.parse(paths, NewGraphQLResponse(paths))
}

func (p *GraphQLParser) parse(paths []string, ret *GraphQLResponse) (*GraphQLResponse, error) {
	rsp := ret
	var err error
	for _, path := range paths {
		rsp, err = p.parsePath(path, rsp)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing graphql")
		}
	}
	return rsp, nil
}

func (p *GraphQLParser) parsePath(path string, ret *GraphQLResponse) (*GraphQLResponse, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	schStr := string(b)
	sch, err := graphql.ParseSchema(schStr, nil)
	if err != nil {
		return nil, err
	}
	intro := sch.Inspect()

	for _, d := range intro.Directives() {
		md := getGraphQLMetadata(d.Description())
		err = ret.Schema.AddOption(&schema.Option{T: "directive", K: d.Name(), V: d.Name(), Metadata: md})
		if err != nil {
			return nil, errors.Wrap(err, "unable to add option")
		}
	}

	for _, t := range intro.Types() {
		err = parseGraphQLType(ret, t, p.logger)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing type from file [" + path + "]")
		}
	}

	return ret, nil
}

func parseScalar(s *schema.Schema, key string, md *schema.Metadata) error {
	if key == "Boolean" || key == "Float" || key == "ID" || key == "Int" || key == "String" {
		return nil
	}
	return s.AddScalar(&schema.Scalar{Key: key, Type: key, Metadata: md})
}

func parseEnum(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	var vals schema.Fields
	evs := t.EnumValues(nil)
	if evs != nil {
		for _, ev := range *evs {
			md := getGraphQLMetadata(ev.Description())
			vals = append(vals, &schema.Field{Key: ev.Name(), Type: schematypes.Wrap(schematypes.EnumValue{}), Metadata: md})
		}
	}
	ret := &schema.Model{Key: key, Type: schema.ModelTypeEnum, Fields: vals, Metadata: md}
	return s.AddModel(ret)
}

func parseInterface(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata, logger logur.Logger) error {
	ret := &schema.Model{Key: key, Type: schema.ModelTypeInterface, Fields: getGraphQLFields(t, logger), Metadata: md}
	return s.AddModel(ret)
}

func parseObject(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata, logger logur.Logger) error {
	ret := &schema.Model{Key: key, Type: schema.ModelTypeStruct, Fields: getGraphQLFields(t, logger), Interfaces: getGraphQLInterfaces(t), Metadata: md}
	return s.AddModel(ret)
}

func parseInputObject(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata, logger logur.Logger) error {
	ret := &schema.Model{Key: key, Type: schema.ModelTypeInput, Fields: getGraphQLFields(t, logger), Metadata: md}
	return s.AddModel(ret)
}

func parseUnion(s *schema.Schema, t *introspection.Type, key string, md *schema.Metadata) error {
	ret := &schema.Model{Key: key, Type: schema.ModelTypeUnion, Metadata: md}
	if t.PossibleTypes() != nil {
		for _, pt := range *t.PossibleTypes() {
			typ := schematypes.Wrap(schematypes.Reference{T: *pt.Name()})
			ret.Fields = append(ret.Fields, &schema.Field{Key: *pt.Name(), Type: typ})
		}
	}
	return s.AddModel(ret)
}
