package task

import (
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/npncore"
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

func (t *Build) Run(project *project.Project, schemata schema.Schemata, options npncore.Entries, logger logur.Logger) Results {
	_, err := util.RunProcessSimple("make build", project.RootPath, logger)
	if err != nil {
		return ErrorResults(t, project, options, err)
	}
	return NewResults(t, project, nil)
}
