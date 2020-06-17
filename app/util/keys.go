package util

import "strings"

const (
	KeyAbout      = "about"
	KeyCreated    = "created"
	KeyContent    = "content"
	KeyDataSource = "dsn"
	KeyEnum       = "enum"
	KeyError      = "error"
	KeyField      = "field"
	KeyID         = "id"
	KeyIdx        = "idx"
	KeyKey        = "key"
	KeyModel      = "model"
	KeyName       = "name"
	KeyProfile    = "profile"
	KeyRole       = "role"
	KeySandbox    = "sandbox"
	KeySchema     = "schema"
	KeyService    = "service"
	KeyStatus     = "status"
	KeyTheme      = "theme"
	KeyTitle      = "title"
	KeyUser       = "user"
)

func Plural(k string) string {
	if len(k) == 0 {
		return k
	}
	switch k {
	case KeySandbox:
		return "sandboxes"
	case KeySchema:
		return "schemata"
	default:
		return k + "s"
	}
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
	switch k {
	case KeyDataSource:
		return "DataSources"
	case KeyID:
		return "IDs"
	case KeyIdx:
		return "Indexes"
	default:
		return Title(Plural(k))
	}
}

func WithID(k string) string {
	return k + "ID"
}
