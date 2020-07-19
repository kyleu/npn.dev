package task

import (
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
)

type Result struct {
	Task     Task             `json:"task"`
	Project  *project.Project `json:"project"`
	Options  util.Entries     `json:"options,omitempty"`
	Data     interface{}      `json:"data,omitempty"`
	Output   []*output.File   `json:"output,omitempty"`
	Duration int              `json:"duration,omitempty"`
}

type AvailableOption struct {
	Key     string `json:"key,omitempty"`
	T       string `json:"t,omitempty"`
	Default string `json:"default,omitempty"`
}
type AvailableOptions []*AvailableOption

type Task interface {
	Key() string
	Title() string
	Description() string
	Options() AvailableOptions
	Run(project *project.Project, schemata schema.Schemata, options util.Entries, logger logur.Logger) (*Result, error)
}

var AllTasks = []Task{&Export{}, &Decorate{}, &Build{}, &Run{}, &Full{}}

func FromString(key string) Task {
	for _, t := range AllTasks {
		if t.Key() == key {
			return t
		}
	}
	return nil
}
