package parser

import (
	"emperror.dev/errors"
	"encoding/xml"
	"github.com/kyleu/npn/app/model/schema"
)

type LiquibaseResponse struct {
	Data []interface{}
}

func NewLiquibaseResponse() *LiquibaseResponse {
	return &LiquibaseResponse{
		Data: make([]interface{}, 0),
	}
}

func (r *LiquibaseResponse) extract(x interface{}, e xml.StartElement, d *xml.Decoder) error {
	err := d.DecodeElement(x, &e)
	if err != nil {
		return errors.Wrap(err, "Error decoding ["+e.Name.Local+"] item")
	}
	r.Data = append(r.Data, x)
	return nil
}

func (r *LiquibaseResponse) Schema() (*schema.Schema, interface{}, error) {
	ret := &schema.Schema{}
	return ret, r, nil
}
