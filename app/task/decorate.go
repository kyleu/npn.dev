package task

import (
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

var KeyDecorate = "decorate"

type Decorate struct {
}

func (t *Decorate) Key() string {
	return KeyDecorate
}

func (t *Decorate) Title() string {
	return "Decorate"
}

func (t *Decorate) Description() string {
	return "Runs language-specific post-processing"
}

func (t *Decorate) Options() AvailableOptions {
	return nil
}

func (t *Decorate) Run(project *project.Project, schemata schema.Schemata, options npncore.Entries, logger logur.Logger) Results {
	_, _, err := npncore.RunProcessSimple("goimports -w .", project.RootPath, logger)
	if err != nil {
		return ErrorResults(t, project, options, err)
	}
	return NewResults(t, project, nil)
}
