package project

import (
	"github.com/kyleu/npn/app/util"
)

type Summary struct {
	Key         string   `json:"key"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	SchemaKeys  []string `json:"schemaKeys,omitempty"`
}

type Summaries []*Summary

type Task struct {
	Key     string       `json:"key"`
	T       string       `json:"t"`
	Options util.Entries `json:"options"`
}
type Tasks []*Task

func (t Tasks) Get(key string) *Task {
	for _, task := range t {
		if task.Key == key {
			return task
		}
	}
	return nil
}

type Project struct {
	Key         string   `json:"key"`
	Title       string   `json:"title,omitempty"`
	SourceURL   string   `json:"sourceURL,omitempty"`
	RootPath    string   `json:"rootPath,omitempty"`
	RootPkg     util.Pkg `json:"rootPkg,omitempty"`
	Description string   `json:"description,omitempty"`
	SchemaKeys  []string `json:"schemaKeys,omitempty"`
	Tasks       Tasks    `json:"tasks,omitempty"`
}
type Projects []*Project

var DefaultProject = &Project{Key: "new"}

func (p *Project) HasSchema(key string) bool {
	for _, s := range p.SchemaKeys {
		if s == key {
			return true
		}
	}
	return false
}
