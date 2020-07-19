package task

import (
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

var KeyFull = "full"

type Full struct {
}

func (t *Full) Key() string {
	return KeyFull
}

func (t *Full) Title() string {
	return "Full"
}

func (t *Full) Description() string {
	return "Runs a bunch of tasks"
}

func (t *Full) Options() AvailableOptions {
	return nil
}

func (t *Full) Run(project *project.Project, schemata schema.Schemata, options util.Entries, logger logur.Logger) (*Result, error) {
	ret := &Result{Task: t, Project: project, Data: []interface{}{}}
	for _, task := range AllTasks {
		if task.Key() == t.Key() {
			continue
		}
		r, err := task.Run(project, schemata, options, logger)
		if err != nil {
			return nil, err
		}
		if len(r.Output) > 0 {
			ret.Output = r.Output
		}
		if r.Data != nil {
			da := ret.Data.([]interface{})
			ret.Data = append(da, r.Data)
		}
	}
	return ret, nil
}
