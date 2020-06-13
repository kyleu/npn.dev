package config

import (
	"github.com/kyleu/npn/app/model/data"
	"github.com/kyleu/npn/app/model/parser"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug    bool
	Parsers  *parser.Parsers
	Files    *data.FileLoader
	Version  string
	Commit   string
	Logger   logur.Logger
}

func (a *AppInfo) Valid() bool {
	return true
}
