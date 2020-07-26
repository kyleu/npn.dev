package task

import (
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

var KeyRun = "run"

type RunProject struct {
}

func (t *RunProject) Key() string {
	return KeyRun
}

func (t *RunProject) Title() string {
	return "RunProject"
}

func (t *RunProject) Description() string {
	return "Runs the project"
}

func (t *RunProject) Options() AvailableOptions {
	return nil
}

func (t *RunProject) Run(project *project.Project, schemata schema.Schemata, options npncore.Entries, logger logur.Logger) Results {
	parsed, err := util.Template("build/{{.Key}}", project)
	if err != nil {
		return ErrorResults(t, project, options, err)
	}
	out, err := util.RunProcessSimple(parsed, project.RootPath, logger)
	if err != nil {
		return ErrorResults(t, project, options, err)
	}
	return NewResults(t, project, nil, out)
}
