package app

import (
	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/imprt"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/app/socket"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice-fs/authfs"
	"github.com/kyleu/npn/npnservice-fs/userfs"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
)

const multiuser = false

type Service struct {
	debug      bool
	files      npncore.FileLoader
	user       user.Service
	auth       auth.Service
	logger     logur.Logger
	Session    *session.Service
	Collection *collection.Service
	Import     *imprt.Service
	Caller     *call.Service
	Socket     *npnconnection.Service
}

var _ npnweb.AppInfo = (*Service)(nil)

func NewService(debug bool, files npncore.FileLoader, redir string, logger logur.Logger) *Service {
	us := userfs.NewServiceFilesystem(multiuser, files, logger)
	sessSvc := session.NewService(files, logger)
	collSvc := collection.NewService(files, logger)
	reqSvc := request.NewService(files, logger)
	callSvc := call.NewService(logger)

	return &Service{
		debug:      debug,
		files:      files,
		user:       us,
		auth:       authfs.NewServiceFS(multiuser, redir, files, logger, us),
		logger:     logger,
		Session:    sessSvc,
		Collection: collSvc,
		Import:     imprt.NewService(files, logger),
		Caller:     callSvc,
		Socket:     socket.NewService(us, sessSvc, collSvc, reqSvc, callSvc, logger),
	}
}

func (c *Service) Debug() bool {
	return c.debug
}

func (c *Service) Files() npncore.FileLoader {
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
