package app

import (
	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/imprt"
	"github.com/kyleu/npn/app/socket"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
)

type Service struct {
	debug      bool
	files      *npncore.FileLoader
	user       *user.Service
	auth       *auth.Service
	version    string
	commit     string
	logger     logur.Logger
	Collection *collection.Service
	Import     *imprt.Service
	Caller     *call.Service
	Socket     *npnconnection.Service
}

func NewService(debug bool, dataDir string, version string, commitHash string, logger logur.Logger) *Service {
	files := npncore.NewFileLoader(dataDir, logger)
	us := user.NewService(files, nil, logger)
	cs := collection.NewService(files, logger)
	return &Service{
		debug:      debug,
		files:      files,
		user:       us,
		auth:       auth.NewService(false, "", nil, logger, us),
		version:    version,
		commit:     commitHash,
		logger:     logger,
		Collection: cs,
		Import:     imprt.NewService(files, logger),
		Caller:     call.NewService(logger),
		Socket:     socket.NewService(cs, logger),
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
