package parsejsonschema

import (
	"fmt"

	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/schema/schematypes"
	"github.com/santhosh-tekuri/jsonschema/v2"
)

func parseType(sch *schema.Schema, js *jsonschema.Schema, pkg []string, modelKey string) schematypes.Wrapped {
	if js.Ref != nil {
		return parseType(sch, js.Ref, pkg, modelKey)
	}

	if len(js.Types) == 1 {
		return parseJsType(sch, js, pkg, modelKey)
	}

	if len(js.AnyOf) > 0 {
		return parseUnion(sch, js.AnyOf, pkg, modelKey)
	}

	return schematypes.NilWrapped
}

func parseJsType(sch *schema.Schema, js *jsonschema.Schema, pkg []string, modelKey string) schematypes.Wrapped {
	switch js.Types[0] {
	case "null":
		return schematypes.ErrorString("null")
	case "boolean":
		return schematypes.Wrap(schematypes.Bool{})
	case "number":
		return schematypes.Wrap(schematypes.Float{})
	case "string":
		pattern := ""
		if js.Pattern != nil {
			pattern = js.Pattern.String()
		}
		return schematypes.Wrap(schematypes.String{Pattern: pattern})
	case "array":
		if js.Items == nil {
			return schematypes.Wrap(schematypes.List{T: schematypes.NilWrapped})
		}
		var is []*jsonschema.Schema
		solo, ok := js.Items.(*jsonschema.Schema)
		if ok {
			is = append(is, solo)
		} else {
			is = js.Items.([]*jsonschema.Schema)
		}
		if len(is) == 1 {
			return parseType(sch, is[0], pkg, modelKey+"Array")
		}
		ret := make([]schematypes.Wrapped, 0, len(is))
		for _, i := range is {
			ret = append(ret, parseType(sch, i, pkg, modelKey))
		}
		if len(ret) == 1 {
			return schematypes.Wrap(schematypes.List{T: ret[0]})
		}
		return schematypes.ErrorString(fmt.Sprintf("unable to process [%v] array types", len(ret)))
	case "object":
		mt, err := parseModel(sch, pkg, modelKey, js)
		if err != nil {
			return schematypes.ErrorWrapped(err)
		}
		return mt

	default:
		return schematypes.ErrorString("invalid js type: " + js.Types[0])
	}
}

func parseUnion(sch *schema.Schema, variants []*jsonschema.Schema, pkg []string, key string) schematypes.Wrapped {
	ret := []schematypes.Wrapped{}
	vIdx := 0
	for _, variant := range variants {
		variantKey := fmt.Sprintf("%vUnion%v", key, vIdx)
		currModel := sch.Models.Get(pkg, key+"Union0")
		for currModel != nil {
			vIdx++
			variantKey = fmt.Sprintf("%vUnion%v", key, vIdx)
			currModel = sch.Models.Get(pkg, variantKey)
		}
		ret = append(ret, parseType(sch, variant, pkg, variantKey))
	}
	switch len(ret) {
	case 0:
		return schematypes.ErrorString("no types for union")
	case 1:
		return ret[0]
	default:
		return schematypes.Wrap(schematypes.Reference{
			Pkg: pkg,
			T:   key + "Union",
		})
	}
}
