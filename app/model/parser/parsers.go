package parser

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"sort"
)

type Parsers struct {
	GraphQL   GraphQLParser
	Protobuf  ProtobufParser
	IntelliJ  IntelliJParser
	Liquibase LiquibaseParser
}


func NewParsers(logger logur.Logger) *Parsers {
	logFor := func(s string) logur.Logger {
		return logur.WithFields(logger, map[string]interface{}{util.KeyService: "intellij"})
	}
	return &Parsers{
		GraphQL: GraphQLParser{Key: schema.OriginGraphQL.Key, logger: logFor(schema.OriginGraphQL.Key)},
		Protobuf: ProtobufParser{Key: schema.OriginProtobuf.Key, logger: logFor(schema.OriginProtobuf.Key)},
		IntelliJ: IntelliJParser{Key: schema.OriginIntelliJ.Key, logger: logFor(schema.OriginIntelliJ.Key)},
		Liquibase: LiquibaseParser{Key: schema.OriginLiquibase.Key, logger: logFor(schema.OriginLiquibase.Key)},
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
	ret := append(gq, append(pb, append(ij, lb...)...)...)
	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].Key < ret[j].Key
	})
	return ret, nil
}

func (p *Parsers) Load(t string, key string) (*schema.Schema, interface{}, error) {
	switch t {
	case p.GraphQL.Key:
		x, err := p.GraphQL.ParseSchemaFile(key)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	case p.Protobuf.Key:
		x, err := p.Protobuf.ParseProtobufFile(key)
		if err != nil {
			return nil, x, err
		}
		return x.Schema, x, nil
	case p.IntelliJ.Key:
		x, err := p.IntelliJ.ParseDataSourceXML(key)
		if err != nil {
			return nil, x, err
		}
		return x.Schema()
	case p.Liquibase.Key:
		x, err := p.Liquibase.ParseChangeLogXML(key)
		if err != nil {
			return nil, x, err
		}
		return x.Schema()
	default:
		return nil, nil, errors.New("invalid type [" + t + "]")
	}
}
