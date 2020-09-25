package npnweb

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"logur.dev/logur"
)

type AppInfo interface {
	Debug() bool
	Files() *npncore.FileLoader
	User() user.Service
	Auth() auth.Service
	Logger() logur.Logger
	Valid() bool
}
