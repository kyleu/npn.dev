package parser

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/parser/graphql"
	parseintellij "github.com/kyleu/npn/app/model/parser/intellij"
	parsejsonschema "github.com/kyleu/npn/app/model/parser/jsonschema"
	"github.com/kyleu/npn/app/model/parser/liquibase"
	"github.com/kyleu/npn/app/model/parser/protobuf"
	"github.com/kyleu/npn/app/model/schema"
	"logur.dev/logur"
	"sort"
)

type Parsers struct {
	GraphQL   *parsegraphql.GraphQLParser
	Protobuf  *parseprotobuf.ProtobufParser
	IntelliJ  *parseintellij.IntelliJParser
	Liquibase *parseliquibase.LiquibaseParser
	JSONSchema *parsejsonschema.JSONSchemaParser
}

func NewParsers(logger logur.Logger) *Parsers {
	return &Parsers{
		GraphQL:   parsegraphql.NewParser(logger),
		Protobuf:  parseprotobuf.NewParser(logger),
		IntelliJ:  parseintellij.NewParser(logger),
		Liquibase: parseliquibase.NewParser(logger),
		JSONSchema: parsejsonschema.NewParser(logger),
	}
}

func (p *Parsers) Detect(root string) ([]schema.DataSource, error){
	gq, err := p.GraphQL.Detect(root)
	if err != nil {
		return nil, err
	}
	pb, err := p.Protobuf.Detect(root)
	if err != nil {
		return nil, err
	}
	ij, err := p.IntelliJ.Detect(root)
	if err != nil {
		return nil, err
	}
	lb, err := p.Liquibase.Detect(root)
	if err != nil {
		return nil, err
	}
	js, err := p.JSONSchema.Detect(root)
	if err != nil {
		return nil, err
	}
	ret := append(gq, append(pb, append(ij, append(lb, js...)...)...)...)
	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].Key < ret[j].Key
	})
	return ret, nil
}

func (p *Parsers) Load(t string, paths []string) (*schema.Schema, interface{}, error) {
	switch t {
	case p.GraphQL.Key:
		x, err := p.GraphQL.ParseSchemaFile(paths)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	case p.Protobuf.Key:
		x, err := p.Protobuf.ParseProtobufFile(paths)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	case p.IntelliJ.Key:
		x, err := p.IntelliJ.ParseDataSourceXML(paths)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	case p.Liquibase.Key:
		x, err := p.Liquibase.ParseChangeLogXML(paths)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	case p.JSONSchema.Key:
		x, err := p.JSONSchema.ParseJSONSchemaFile(paths)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	default:
		return nil, nil, errors.New("invalid parser type [" + t + "]")
	}
}

func (p *Parsers) Refresh(sch *schema.Schema) (*schema.Schema, error) {
	newSchema, _, err :=  p.Load(sch.Metadata.Origin.Key, sch.Paths)
	if err != nil {
		return nil, err
	}
	newSchema.Key = sch.Key
	newSchema.Title = sch.Title
	return newSchema, nil
}
