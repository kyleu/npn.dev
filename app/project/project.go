package project

import (
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/npncore"
)

type TaskDefinition struct {
	Key     string          `json:"key"`
	T       string          `json:"t"`
	Options npncore.Entries `json:"options"`
}

func (d *TaskDefinition) Clone() *TaskDefinition {
	return &TaskDefinition{Key: d.Key, T: d.T, Options: d.Options.Clone()}
}

type TaskDefinitions []*TaskDefinition

func (t TaskDefinitions) Get(key string) *TaskDefinition {
	for _, task := range t {
		if task.Key == key {
			return task
		}
	}
	return nil
}

func (t TaskDefinitions) Replacing(origKey string, dst *TaskDefinition) TaskDefinitions {
	ret := make(TaskDefinitions, 0, len(t))
	matched := false
	for _, orig := range t {
		if orig.Key == origKey {
			matched = true
			ret = append(ret, dst)
		} else {
			ret = append(ret, orig)
		}
	}
	if !matched {
		ret = append(ret, dst)
	}
	return ret
}

func (t TaskDefinitions) Without(key string) TaskDefinitions {
	ret := make(TaskDefinitions, 0, len(t))
	for _, orig := range t {
		if orig.Key != key {
			ret = append(ret, orig)
		}
	}
	return ret
}

type Project struct {
	Key         string                 `json:"key"`
	Title       string                 `json:"title,omitempty"`
	Org         string                 `json:"org,omitempty"`
	SourceURL   string                 `json:"sourceURL,omitempty"`
	RootPath    string                 `json:"rootPath,omitempty"`
	RootPkg     util.Pkg               `json:"rootPkg,omitempty"`
	Description string                 `json:"description,omitempty"`
	Prototype   string                 `json:"prototype,omitempty"`
	SchemaKeys  []string               `json:"schemaKeys,omitempty"`
	Models      ModelRefs              `json:"models,omitempty"`
	Options     map[string]interface{} `json:"options,omitempty"`
	Tasks       TaskDefinitions        `json:"tasks,omitempty"`
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

func (p *Project) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, len(p.Options)+7)
	ret["Key"] = p.Key
	ret["Title"] = p.Title
	ret["SourceURL"] = p.SourceURL
	ret["RootPath"] = p.RootPath
	ret["RootPkg"] = p.RootPkg
	ret["Description"] = p.Description
	ret["Prototype"] = p.Prototype
	for k, v := range p.Options {
		ret[k] = v
	}
	return ret
}
