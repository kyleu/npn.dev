package parsejsonschema

import (
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"github.com/kyleu/npn/app/util"
	"strings"
	"text/scanner"
)

func getJSONSchemaMetadata(pos scanner.Position, comments ...string) *schema.Metadata {
	cmts := make([]string, 0, len(comments))
	for _, cs := range comments {
		l := strings.TrimSpace(cs)
		cmts = append(cmts, l)
	}
	return &schema.Metadata{
		Comments: cmts,
		Origin:   schema.OriginJSONSchema,
		Source:   util.FilenameOf(pos.Filename),
		Line:     pos.Line,
		Column:   pos.Column - 1,
	}
}

func getJSONSchemaType(currPkg []string, t string, optional bool, repeated bool) schematypes.Wrapped {
	var ret schematypes.Type
	if optional {
		ret = schematypes.Option{T: getJSONSchemaType(currPkg, t, false, repeated)}
	} else if repeated {
		ret = schematypes.List{T: getJSONSchemaType(currPkg, t, optional, false)}
	} else {
		ret = getTypeForJSONSchemaName(currPkg, t)
	}
	return schematypes.Wrap(ret)
}

func getTypeForJSONSchemaName(currPkg []string, name string) schematypes.Wrapped {
	var ret schematypes.Type
	switch name {
	case "Boolean", "Bool":
		ret = schematypes.Bool{}
	case "int64", "int32":
		ret = schematypes.Int{}
	case "Float":
		ret = schematypes.Float{}
	case "string":
		ret = schematypes.String{}
	default:
		pkg, key := util.SplitPackage(name)
		if len(pkg) == 0 {
			ret = schematypes.Reference{Pkg: currPkg, T: key}
		} else {
			ret = schematypes.Reference{Pkg: pkg, T: key}
		}
	}
	return schematypes.Wrap(ret)
}
