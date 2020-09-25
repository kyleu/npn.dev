package app

import (
	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/imprt"
	"github.com/kyleu/npn/app/socket"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice-fs/userfs"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice-fs/authfs"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
)

type Service struct {
	debug      bool
	files      *npncore.FileLoader
	user       user.Service
	auth       auth.Service
	logger     logur.Logger
	Collection *collection.Service
	Import     *imprt.Service
	Caller     *call.Service
	Socket     *npnconnection.Service
}

func NewService(debug bool, dataDir string, logger logur.Logger) *Service {
	files := npncore.NewFileLoader(dataDir, logger)
	us := userfs.NewServiceFilesystem(false, files, logger)
	collSvc := collection.NewService(files, logger)
	callSvc := call.NewService(logger)
	return &Service{
		debug:      debug,
		files:      files,
		user:       us,
		// auth:       authdb.NewServiceDatabase(false, "", nil, logger, us),
		auth:       authfs.NewServiceNoop(),
		logger:     logger,
		Collection: collSvc,
		Import:     imprt.NewService(files, logger),
		Caller:     callSvc,
		Socket:     socket.NewService(collSvc, callSvc, logger),
	}
}

func (c *Service) Debug() bool {
	return c.debug
}

func (c *Service) Files() *npncore.FileLoader {
	return c.files
}

func (c *Service) User() user.Service {
	return c.user
}

func (c *Service) Auth() auth.Service {
	return c.auth
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
