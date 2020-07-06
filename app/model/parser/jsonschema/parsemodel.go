package parsejsonschema

import (
	"sort"
	"strings"

	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"github.com/santhosh-tekuri/jsonschema/v2"
)

func parseModel(sch *schema.Schema, pkg util.Pkg, key string, js *jsonschema.Schema) (schematypes.Wrapped, error) {
	ptr := strings.TrimPrefix(strings.TrimPrefix(js.Ptr, "#"), "/")
	if len(ptr) > 0 {
		pkg, key = util.SplitPackageSlash(ptr)
	}

	fieldNames := make([]string, 0, len(js.Properties))
	for fieldName := range js.Properties {
		fieldNames = append(fieldNames, fieldName)
	}
	sort.Strings(fieldNames)

	model := &schema.Model{Key: key, Pkg: pkg, Type: schema.ModelTypeStruct, Description: js.Description, Metadata: nil}

	for _, fieldName := range fieldNames {
		isRequired := false
		for _, s := range js.Required {
			if fieldName == s {
				isRequired = true
			}
		}
		parseField(sch, model, fieldName, js.Properties[fieldName], isRequired)
	}

	currModel := sch.Models.Get(pkg, key)
	if currModel == nil {
		err := sch.AddModel(model)
		if err != nil {
			return schematypes.ErrorWrapped(err), err
		}
	}

	return schematypes.ReferenceWrapped(pkg, key), nil
}
