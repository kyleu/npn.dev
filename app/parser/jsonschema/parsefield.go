package parsejsonschema

import (
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/schema/schematypes"
	"github.com/santhosh-tekuri/jsonschema/v2"
)

func parseField(sch *schema.Schema, model *schema.Model, key string, js *jsonschema.Schema, isRequired bool) *schema.Field {
	t := parseType(sch, js, model.Pkg, key)
	if !isRequired {
		t = schematypes.OptionWrapped(t)
	}
	md := &schema.Metadata{Description: js.Description, Origin: schema.OriginJSONSchema}
	ret := &schema.Field{Key: key, Type: t, Metadata: md}
	err := model.AddField(ret)
	if err != nil {
		ret.Type = schematypes.ErrorWrapped(err)
	}
	return ret
}
