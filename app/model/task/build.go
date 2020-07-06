package task

import (
	"github.com/kyleu/npn/app/model/project"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

var KeyBuild = "build"

type Build struct {
}

func (t *Build) Key() string {
	return KeyBuild
}

func (t *Build) Title() string {
	return "Build"
}

func (t *Build) Description() string {
	return "Builds the project"
}

func (t *Build) Options() AvailableOptions {
	return nil
}

func (t *Build) Run(project *project.Project, schemata schema.Schemata, options util.Entries, logger logur.Logger) (*Result, error) {
	out, err := util.RunProcessSimple("make build", project.RootPath, logger)
	if err != nil {
		return nil, err
	}
	return &Result{Task: t, Project: project, Data: out}, nil
}
