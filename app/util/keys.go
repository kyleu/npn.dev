package util

import "strings"

const (
	KeyAbout      = "about"
	KeyCreated    = "created"
	KeyContent    = "content"
	KeyEnum       = "enum"
	KeyError      = "error"
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
	KeyUnion      = "union"
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
	case KeyID:
		return "ID"
	case KeyIdx:
		return "Index"
	}
	return strings.ToUpper(k[0:1]) + k[1:]
}

func PluralTitle(k string) string {
	return Title(Plural(k))
}

func WithID(k string) string {
	return k + "ID"
}
