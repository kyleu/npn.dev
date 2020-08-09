package task

import (
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
	"time"
)

type Result struct {
	Task     Task             `json:"task"`
	Project  *project.Project `json:"project"`
	Options  npncore.Entries  `json:"options,omitempty"`
	Data     npncore.Data     `json:"data,omitempty"`
	Messages []string         `json:"messages,omitempty"`
	Duration int              `json:"duration,omitempty"`
}

func NewResult(task Task, project *project.Project, options npncore.Entries, msgs ...string) *Result {
	return &Result{Task: task, Project: project, Options: options, Data: make(npncore.Data), Messages: msgs}
}

type Results []*Result

func NewResults(task Task, project *project.Project, options npncore.Entries, msgs ...string) Results {
	return Results{NewResult(task, project, options, msgs...)}
}

func ErrorResults(task Task, project *project.Project, options npncore.Entries, err error) Results {
	msgs := []string{"Error: " + err.Error()}
	d := npncore.Data{"data": err.Error()}
	res := Result{Task: task, Project: project, Options: options, Data: d, Messages: msgs}
	return Results{&res}
}

type AvailableOption struct {
	Key     string `json:"key,omitempty"`
	T       string `json:"t,omitempty"`
	Default string `json:"default,omitempty"`
	Desc    string `json:"desc,omitempty"`
}
type AvailableOptions []*AvailableOption

type Task interface {
	Key() string
	Title() string
	Description() string
	Options() AvailableOptions
	Run(project *project.Project, schemata schema.Schemata, options npncore.Entries, logger logur.Logger) Results
}

var AllTasks = []Task{&Bootstrap{}, &Export{}, &Decorate{}, &Build{}, &RunProject{}}

func FromString(key string) Task {
	for _, t := range AllTasks {
		if t.Key() == key {
			return t
		}
	}
	return nil
}

func RunTask(proj *project.Project, schemata schema.Schemata, t Task, options npncore.Entries, logger logur.Logger) Results {
	startNanos := time.Now().UnixNano()
	r := t.Run(proj, schemata, options, logger)
	if len(r) > 0 {
		delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
		r[0].Duration = int(delta)
	}
	return r
}
