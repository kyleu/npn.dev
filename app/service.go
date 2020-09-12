package app

import (
	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/socket"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
)

type Service struct {
	debug   bool
	files   *npncore.FileLoader
	user    *user.Service
	auth    *auth.Service
	version string
	commit  string
	logger  logur.Logger
	Caller  *call.Service
	Socket  *npnconnection.Service
}

func NewService(debug bool, version string, commitHash string, logger logur.Logger) *Service {
	files := npncore.NewFileLoader(".", logger)
	us := user.NewService(files, nil, logger)
	au := auth.NewService(false, "", nil, logger, us)
	call := call.NewService(logger)
	sock := socket.NewService(logger)
	return &Service{
		debug:   debug,
		files:   files,
		user:    us,
		auth:    au,
		version: version,
		commit:  commitHash,
		logger:  logger,
		Caller:  call,
		Socket:  sock,
	}
}

func (c *Service) Debug() bool {
	return c.debug
}

func (c *Service) Files() *npncore.FileLoader {
	return c.files
}

func (c *Service) User() *user.Service {
	return c.user
}

func (c *Service) Auth() *auth.Service {
	return c.auth
}

func (c *Service) Version() string {
	return c.version
}

func (c *Service) Commit() string {
	return c.commit
}

func (c *Service) Logger() logur.Logger {
	return c.logger
}

func (c *Service) Valid() bool {
	return true
}

func Svc(a npnweb.AppInfo) *Service {
	return a.(*Service)
}
