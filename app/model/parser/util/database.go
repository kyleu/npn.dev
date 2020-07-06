package parseutil

import (
	"strconv"
	"strings"

	"github.com/kyleu/npn/app/util"

	"github.com/kyleu/npn/app/model/schema/schematypes"
)

func ParseDatabaseType(pkg util.Pkg, t string, optional bool) schematypes.Wrapped {
	if optional {
		return schematypes.OptionWrapped(ParseDatabaseType(pkg, t, false))
	}

	tIndex := strings.Index(t, "|")
	if tIndex > -1 {
		t = t[0:tIndex]
	}

	args := ""
	parenIndex := strings.Index(t, "(")
	if parenIndex > -1 {
		args = t[parenIndex+1 : len(t)-1]
		t = t[0:parenIndex]
	}
	bracketIndex := strings.Index(t, "[]")
	if bracketIndex > -1 {
		return schematypes.Wrap(schematypes.List{T: ParseDatabaseType(pkg, t[0:bracketIndex], false)})
	}

	var ret schematypes.Type
	lowered := strings.ToLower(t)
	switch lowered {
	case "bool", "boolean":
		ret = schematypes.Bool{}
	case "b", "byte":
		ret = schematypes.Byte{}
	case "char":
		ret = schematypes.Char{}
	case "date":
		ret = schematypes.Date{}
	case "money", "real", "double precision", "numeric":
		ret = schematypes.Float{}
	case "int", "integer", "smallint", "bigint":
		ret = schematypes.Int{}
	case "json":
		ret = schematypes.JSON{}
	case "text", "varchar", "_varchar", "bpchar":
		maxLength := int64(0)
		if len(args) > 0 {
			maxLength, _ = strconv.ParseInt(args, 10, 64)
		}
		ret = schematypes.String{MaxLength: int(maxLength)}
	case "time", "time without time zone", "time with time zone":
		ret = schematypes.Time{}
	case "timestamp", "timestamp without time zone":
		ret = schematypes.Timestamp{}
	case "timestamp with time zone":
		ret = schematypes.TimestampZoned{}
	case "uuid":
		ret = schematypes.UUID{}
	case "xml":
		ret = schematypes.XML{}

	case "bit", "bit varying":
		ret = schematypes.List{T: schematypes.Wrap(schematypes.Bit{})}
	case "oid":
		ret = schematypes.Int{}
	case "bytea":
		ret = schematypes.List{T: schematypes.Wrap(schematypes.Byte{})}
	case "hstore":
		ret = schematypes.Map{K: schematypes.StringWrapped, V: schematypes.StringWrapped}

	default:
		ret = schematypes.Reference{Pkg: pkg, T: t}
	}

	return schematypes.Wrap(ret)
}
