package parseprotobuf

import (
	"fmt"
	"github.com/emicklei/proto"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"github.com/kyleu/npn/app/util"
	"strings"
	"text/scanner"
)

func getProtobufMetadata(pos scanner.Position, comments ...*proto.Comment) *schema.Metadata {
	cmt := make([]string, 0)
	for _, cs := range comments {
		if cs != nil {
			for _, l := range cs.Lines {
				l = strings.TrimSpace(l)
				cmt = append(cmt, l)
			}
		}
	}
	return &schema.Metadata{
		Comments: cmt,
		Origin:   schema.OriginProtobuf,
		Source:   util.FilenameOf(pos.Filename),
		Line:     pos.Line,
		Column:   pos.Column - 1,
	}
}

func getProtobufType(currPkg []string, t string, optional bool, repeated bool, options []*proto.Option) schematypes.Wrapped {
	var ret schematypes.Type
	for _, opt := range options {
		panic(fmt.Sprintf("option [%v] provided for proto type", opt))
	}
	if optional {
		ret = schematypes.Option{T: getProtobufType(currPkg, t, false, repeated, options)}
	} else if repeated {
		ret = schematypes.List{T: getProtobufType(currPkg, t, optional, false, options)}
	} else {
		ret = getTypeForProtobufName(currPkg, t)
	}
	return schematypes.Wrap(ret)
}

func getTypeForProtobufName(currPkg []string, name string) schematypes.Wrapped {
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
