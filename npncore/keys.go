package npncore

import (
	"strings"

	"github.com/jinzhu/inflection"
)

const (
	KeyAbout      = "about"
	KeyAct        = "act"
	KeyAction     = "action"
	KeyAdmin      = "admin"
	KeyAuth       = "auth"
	KeyCategory   = "category"
	KeyChoice     = "choice"
	KeyCollection = "collection"
	KeyComment    = "comment"
	KeyConnection = "connection"
	KeyContent    = "content"
	KeyCreated    = "created"
	KeyData       = "data"
	KeyDetail     = "detail"
	KeyEmail      = "email"
	KeyError      = "error"
	KeyExport     = "export"
	KeyFile       = "file"
	KeyFmt        = "fmt"
	KeyGraphQL    = "graphql"
	KeyGraphiQL   = "graphiql"
	KeyHistory    = "history"
	KeyHTML       = "html"
	KeyID         = "id"
	KeyIdx        = "idx"
	KeyInvitation = "invitation"
	KeyJSON       = "json"
	KeyKey        = "key"
	KeyMember     = "member"
	KeyMigration  = "migration"
	KeyModel      = "model"
	KeyModules    = "modules"
	KeyName       = "name"
	KeyNote       = "note"
	KeyOwner      = "owner"
	KeyPermission = "permission"
	KeyProfile    = "profile"
	KeyProvider   = "provider"
	KeyQuery      = "query"
	KeyReport     = "report"
	KeyRequest    = "request"
	KeyRole       = "role"
	KeyRoutes     = "routes"
	KeySandbox    = "sandbox"
	KeyService    = "service"
	KeySession    = "session"
	KeySlug       = "slug"
	KeySocket     = "socket"
	KeySQL        = "sql"
	KeyStatus     = "status"
	KeySvc        = "svc"
	KeySystem     = "system"
	KeyTest       = "test"
	KeyTheme      = "theme"
	KeyTitle      = "title"
	KeyTranscript = "transcript"
	KeyUser       = "user"
	KeyVoyager    = "voyager"
)

func Plural(s string) string {
	return inflection.Plural(s)
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
	return Plural(Title(k))
}

func WithID(k string) string {
	return k + "ID"
}

func WithDBID(k string) string {
	return k + "_id"
}
