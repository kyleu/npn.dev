package parser

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

func getProtobufType(t string, optional bool, repeated bool, options []*proto.Option) schematypes.Wrapped {
	var ret schematypes.Type
	for _, opt := range options {
		println(fmt.Sprintf("option [%v] provided for proto type", opt))
	}
	if optional {
		ret = schematypes.Option{T: getProtobufType(t, false, repeated, options)}
	} else if repeated {
		ret = schematypes.List{T: getProtobufType(t, optional, false, options)}
	} else {
		ret = getTypeForProtobufName(t)
	}
	return schematypes.Wrap(ret)
}

func getTypeForProtobufName(name string) schematypes.Wrapped {
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
		ret = schematypes.Reference{T: name}
	}
	return schematypes.Wrap(ret)
}
