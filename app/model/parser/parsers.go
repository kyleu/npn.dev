package parser

import (
	"sort"

	parsegraphql "github.com/kyleu/npn/app/model/parser/graphql"
	parseintellij "github.com/kyleu/npn/app/model/parser/intellij"
	parsejsonschema "github.com/kyleu/npn/app/model/parser/jsonschema"
	parseprotobuf "github.com/kyleu/npn/app/model/parser/protobuf"
	parseutil "github.com/kyleu/npn/app/model/parser/util"
	"github.com/kyleu/npn/app/model/schema"
	"logur.dev/logur"
)

type Parser interface {
	Type() schema.Origin
	Detect(root string) ([]schema.DataSource, error)
	IsValid(firstChars string) error
	Parse(paths []string) (*parseutil.ParseResponse, error)
}

type Parsers struct {
	GraphQL    *parsegraphql.GraphQLParser
	Protobuf   *parseprotobuf.ProtobufParser
	IntelliJ   *parseintellij.IntelliJParser
	JSONSchema *parsejsonschema.JSONSchemaParser
}

func NewParsers(logger logur.Logger) *Parsers {
	return &Parsers{
		GraphQL:    parsegraphql.NewParser(logger),
		Protobuf:   parseprotobuf.NewParser(logger),
		IntelliJ:   parseintellij.NewParser(logger),
		JSONSchema: parsejsonschema.NewParser(logger),
	}
}

func (p *Parsers) Detect(root string) ([]schema.DataSource, error) {
	ret := make([]schema.DataSource, 0)
	for _, parser := range p.all() {
		x, err := parser.Detect(root)
		if err != nil {
			return nil, err
		}
		ret = append(ret, x...)
	}
	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].Key < ret[j].Key
	})
	return ret, nil
}

func (p *Parsers) IsValid(path string) (schema.Origin, error) {
	firstK, err := parseutil.ReadFirstK(path)
	if err != nil {
		return schema.OriginUnknown, err
	}
	for _, parser := range p.all() {
		v := parser.IsValid(firstK)
		if v == nil {
			return parser.Type(), nil
		}
	}
	return schema.OriginUnknown, nil
}

func (p *Parsers) all() []Parser {
	return []Parser{p.GraphQL, p.Protobuf, p.IntelliJ, p.JSONSchema}
}

func (p *Parsers) fromType(t schema.Origin) Parser {
	for _, parser := range p.all() {
		if t == parser.Type() {
			return parser
		}
	}
	return nil
}

func (p *Parsers) Load(t schema.Origin, paths []string) (*schema.Schema, interface{}, error) {
	x, err := p.fromType(t).Parse(paths)
	if err != nil {
		return nil, x, err
	}
	return x.Schema, x, nil
}

func (p *Parsers) Refresh(sch *schema.Schema) (*schema.Schema, error) {
	t := sch.Metadata.Origin
	newSchema, _, err := p.Load(t, sch.Paths)
	if err != nil {
		return nil, err
	}
	newSchema.Key = sch.Key
	newSchema.Title = sch.Title
	return newSchema, nil
}
