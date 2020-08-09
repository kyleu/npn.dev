package task

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/export"
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

var KeyExport = "export"

type Export struct {
}

func (t *Export) Key() string {
	return KeyExport
}

func (t *Export) Title() string {
	return "Export"
}

func (t *Export) Description() string {
	return "Generates code for the project"
}

func (t *Export) Options() AvailableOptions {
	return AvailableOptions{
		{Key: "schema", T: "schema"},
		{Key: "include", T: "models", Default: "*"},
		{Key: "exclude", T: "models"},
		{Key: "package", T: "string", Desc: "package override for all exported models"},
	}
}

func (t *Export) Run(project *project.Project, schemata schema.Schemata, options npncore.Entries, logger logur.Logger) Results {
	var ret []*output.File
	nr := output.GoNameRegistry()
	schemaOpt := options.GetStringArray("schema")
	pkgOpt := util.PkgFromString(options.GetString("package"))
	for _, sch := range schemata {
		if len(schemaOpt) == 0 || npncore.StringArrayContains(schemaOpt, sch.Key) {
			for _, model := range sch.Models {
				p := model.Pkg
				if len(pkgOpt) > 0 {
					p = pkgOpt
				}
				file := output.NewGoFile(p, model.Key)
				export.WriteGo(file, model, nr)
				ret = append(ret, file)
			}
		}
	}

	res := NewResult(t, project, nil)

	x, err := res.applyOutput(ret, project.RootPath)
	if err != nil {
		return ErrorResults(t, project, options, errors.Wrap(err, "error applying output"))
	}
	res.Data["output"] = x

	return Results{res}
}
