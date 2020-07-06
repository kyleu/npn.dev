package task

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/export"
	"github.com/kyleu/npn/app/model/output"
	"github.com/kyleu/npn/app/model/project"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
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
		{Key: "include", T: "string", Default: "*"},
	}
}

func (t *Export) Run(project *project.Project, schemata schema.Schemata, options util.Entries, logger logur.Logger) (*Result, error) {
	var ret []*output.File
	nr := output.GoNameRegistry()
	for _, sch := range schemata {
		for _, model := range sch.Models {
			file := output.NewGoFile(project, model.Pkg, model.Key)
			export.WriteGo(file, model, nr)
			ret = append(ret, file)
		}
	}

	res := &Result{Task: t, Project: project, Output: ret}

	x, err := res.applyOutput(ret, project.RootPath)
	if err != nil {
		return nil, errors.Wrap(err, "error applying output")
	}
	res.Data = x

	return res, nil
}
