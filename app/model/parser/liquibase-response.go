package parser

import (
	"emperror.dev/errors"
	"encoding/xml"
	"fmt"
	"github.com/kyleu/npn/app/model/schema"
)

type LiquibaseResponse struct {
	RootFile string `json:"root"`
	Data     []interface{}
	Schema   *schema.Schema `json:"schema"`
}

func NewLiquibaseResponse(paths []string) *LiquibaseResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginLiquibase, Source: paths[0]}
	return &LiquibaseResponse{
		RootFile: paths[0],
		Data:     make([]interface{}, 0),
		Schema:   schema.NewSchema(paths[0], paths, &md),
	}
}

func (r *LiquibaseResponse) extract(x interface{}, e xml.StartElement, d *xml.Decoder) error {
	err := d.DecodeElement(x, &e)
	if err != nil {
		return errors.Wrap(err, "Error decoding ["+e.Name.Local+"] item")
	}
	r.Data = append(r.Data, x)
	return r.process(x)
}

func (r *LiquibaseResponse) process(x interface{}) error {
	var err error
	switch msg := x.(type) {
	case *lCreateTable:
		err = r.Schema.AddModel(&schema.Model{
			Key:    msg.Name,
			Type:   schema.ModelTypeStruct,
			Fields: r.toFields(msg.Columns),
		})
	case *lAddForeignKeyConstraint:
	case *lAddUniqueConstraint:
	case *lCreateIndex:
	default:
		err = errors.New(fmt.Sprintf("invalid liquibase message [%T]", x))
	}
	return err
}

func (r *LiquibaseResponse) toFields(columns []lColumn) schema.Fields {
	ret := make(schema.Fields, 0, len(columns))
	for _, col := range columns {
		ret = append(ret, &schema.Field{
			Key:  col.Name,
			Type: parseDatabaseType(col.T),
		})
	}
	return ret
}
