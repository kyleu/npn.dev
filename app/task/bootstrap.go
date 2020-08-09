package task

import (
	"fmt"
	"os"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/bootstrap"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

var KeyBootstrap = "bootstrap"

type Bootstrap struct {
}

func (t *Bootstrap) Key() string {
	return KeyBootstrap
}

func (t *Bootstrap) Title() string {
	return "Bootstrap"
}

func (t *Bootstrap) Description() string {
	return "Generates code for the project"
}

func (t *Bootstrap) Options() AvailableOptions {
	return AvailableOptions{}
}

func (t *Bootstrap) Run(project *project.Project, schemata schema.Schemata, options npncore.Entries, logger logur.Logger) Results {
	msg := "Skipped bootstrap, project exists"
	f, err := os.Open(project.RootPath)
	if err != nil || f == nil {
		proto := bootstrap.PrototypeFromString(project.Prototype)
		err = bootstrap.Extract(proto, project, logger)
		if err != nil {
			return ErrorResults(t, project, options, errors.Wrap(err, "can't extract ["+proto.Key+"]"))
		}
		_, _, err = npncore.RunProcessSimple("git init", project.RootPath, logger)
		if err != nil {
			return ErrorResults(t, project, options, err)
		}
		_, _, err = npncore.RunProcessSimple("git add .", project.RootPath, logger)
		if err != nil {
			return ErrorResults(t, project, options, err)
		}
		_, _, err = npncore.RunProcessSimple("git commit -m initial_commit", project.RootPath, logger)
		if err != nil {
			return ErrorResults(t, project, options, err)
		}

		msg = fmt.Sprintf("Extracted prototype [%v] to [%v]", proto.Key, project.RootPath)
	}

	return NewResults(t, project, nil, msg)
}
