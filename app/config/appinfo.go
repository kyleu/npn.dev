package config

import (
	"github.com/kyleu/npn/app/parser"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/npncore"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug    bool
	Parsers  *parser.Parsers
	Files    *npncore.FileLoader
	Schemata *schema.Service
	Projects *project.Service
	Version  string
	Commit   string
	Logger   logur.Logger
}

func (a *AppInfo) Valid() bool {
	return true
}
