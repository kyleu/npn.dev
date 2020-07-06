package task

import (
	"github.com/kyleu/npn/app/model/project"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

var KeyRun = "run"

type Run struct {
}

func (t *Run) Key() string {
	return KeyRun
}

func (t *Run) Title() string {
	return "Run"
}

func (t *Run) Description() string {
	return "Runs the project"
}

func (t *Run) Options() AvailableOptions {
	return nil
}

func (t *Run) Run(project *project.Project, schemata schema.Schemata, options util.Entries, logger logur.Logger) (*Result, error) {
	parsed, err := util.Template("build/{{.Key}}", project)
	if err != nil {
		return nil, err
	}
	out, err := util.RunProcessSimple(parsed, project.RootPath, logger)
	if err != nil {
		return nil, err
	}
	return &Result{Task: t, Project: project, Data: out}, nil
}
