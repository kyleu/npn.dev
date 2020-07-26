package config

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/parser"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/task"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug    bool
	Parsers  *parser.Parsers
	Files    *npncore.FileLoader
	Schemata *schema.Service
	Projects *project.Service
	Version  string
	Commit   string
	Logger   logur.Logger
}

func (a *AppInfo) Valid() bool {
	return true
}

func (a *AppInfo) RunTask(t task.Task, projectKey string, options npncore.Entries) task.Results {
	proj, err := a.Projects.Load(projectKey)
	if err != nil {
		err = errors.Wrap(err, "cannot load project ["+projectKey+"]")
		return task.ErrorResults(t, proj, options, err)
	}
	var schemata schema.Schemata
	for _, schemaKey := range proj.SchemaKeys {
		sch, err := a.Schemata.Load(schemaKey)
		if err != nil {
			err = errors.Wrap(err, "cannot load schema ["+schemaKey+"] for project ["+proj.Key+"]")
			return task.ErrorResults(t, proj, options, err)
		}
		schemata = append(schemata, sch)
	}
	return task.RunTask(proj, schemata, t, options, a.Logger)
}
