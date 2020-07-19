package output

import (
	"github.com/kyleu/npn/app/util"
)

var defaultGoRegistry = util.NewNameRegistry(map[string]*util.Name{
	"date":      {Pkg: util.Pkg{"time"}, Name: "Time"},
	"json":      {Pkg: util.Pkg{"json"}, Name: "RawMessage"},
	"timestamp": {Pkg: util.Pkg{"time"}, Name: "Time"},
	"uuid":      {Pkg: util.Pkg{"uuid"}, Name: "UUID"},
}, map[string]string{"Csv": "CSV", "Id": "ID", "Json": "JSON", "Sql": "SQL", "Xml": "XML"})

func GoNameRegistry() *util.NameRegistry {
	r := *defaultGoRegistry
	return &r
}
