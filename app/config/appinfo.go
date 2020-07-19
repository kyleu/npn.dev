package config

import (
	"github.com/kyleu/npn/app/util"
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/parser"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/task"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug    bool
	Parsers  *parser.Parsers
	Files    *util.FileLoader
	Schemata *schema.Cache
	Projects *project.Cache
	Version  string
	Commit   string
	Logger   logur.Logger
}

func (a *AppInfo) Valid() bool {
	return true
}

func (a *AppInfo) RunTask(t task.Task, projectKey string, options util.Entries) (*task.Result, error) {
	proj, err := a.Projects.Load(projectKey)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load project ["+projectKey+"]")
	}
	var schemata schema.Schemata
	for _, schemaKey := range proj.SchemaKeys {
		sch, err := a.Schemata.Load(schemaKey)
		if err != nil {
			return nil, errors.Wrap(err, "cannot load schema ["+schemaKey+"] for project ["+proj.Key+"]")
		}
		schemata = append(schemata, sch)
	}
	startNanos := time.Now().UnixNano()
	r, err := t.Run(proj, schemata, options, a.Logger)
	if r != nil {
		delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
		r.Duration = int(delta)
	}
	return r, err
}
