package task

import (
	"github.com/kyleu/npn/app/model/project"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
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

func (t *Decorate) Run(project *project.Project, schemata schema.Schemata, options util.Entries, logger logur.Logger) (*Result, error) {
	out, err := util.RunProcessSimple("goimports -w .", project.RootPath, logger)
	if err != nil {
		return nil, err
	}
	return &Result{Task: t, Project: project, Data: out}, nil
}
