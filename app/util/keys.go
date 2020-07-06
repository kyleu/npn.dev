package util

import (
	"github.com/jinzhu/inflection"
	"strings"
)

const (
	KeyAbout      = "about"
	KeyCreated    = "created"
	KeyContent    = "content"
	KeyDataSource = "dsn"
	KeyError      = "error"
	KeyFile       = "file"
	KeyID         = "id"
	KeyIdx        = "idx"
	KeyKey        = "key"
	KeyModel      = "model"
	KeyModules    = "modules"
	KeyName       = "name"
	KeyProfile    = "profile"
	KeyProject    = "project"
	KeyRole       = "role"
	KeyRoutes     = "routes"
	KeySandbox    = "sandbox"
	KeySchema     = "schema"
	KeyService    = "service"
	KeyStatus     = "status"
	KeyTask       = "task"
	KeyTheme      = "theme"
	KeyTitle      = "title"
	KeyUser       = "user"
)

func Plural(s string) string {
	return inflection.Plural(s)
}

func Title(k string) string {
	if len(k) == 0 {
		return k
	}
	switch k {
	case KeyDataSource:
		return "DataSource"
	case KeyID:
		return "ID"
	case KeyIdx:
		return "Index"
	}
	return strings.ToUpper(k[0:1]) + k[1:]
}

func PluralTitle(k string) string {
	return Plural(Title(k))
}
