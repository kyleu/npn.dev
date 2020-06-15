package parser

import (
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"strings"
)

func parseDatabaseType(t string) schematypes.Wrapped {
	var ret schematypes.Type
	args := ""
	parenIndex := strings.Index(t, "(")
	if parenIndex > -1 {
		args = t[parenIndex:]
		t = t[0:parenIndex]
	}
	lowered := strings.ToLower(t)
	switch lowered {
	case "date":
		ret = schematypes.Unknown{X: "TODO: " + t}
	case "int", "integer":
		ret = schematypes.Int{}
	case "json":
		ret = schematypes.Unknown{X: "TODO: " + t}
	case "text", "varchar":
		ret = schematypes.String{}
	case "timestamp", "timestamp without time zone":
		ret = schematypes.Unknown{X: "TODO: " + t}
	case "uuid":
		ret = schematypes.Unknown{X: "TODO: " + t}
	default:
		ret = schematypes.Error{Message: "invalid database kind [" + t + args + "]"}
	}
	return schematypes.Wrap(ret)
}
