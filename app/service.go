package app

import (
	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnservice-fs/authfs"
	"github.com/kyleu/libnpn/npnservice-fs/userfs"
	"github.com/kyleu/libnpn/npnservice/auth"
	"github.com/kyleu/libnpn/npnservice/user"
	"github.com/kyleu/libnpn/npnweb"
	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/imprt"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/search"
	"github.com/kyleu/npn/app/session"
	"github.com/kyleu/npn/app/socket"
	"github.com/sirupsen/logrus"
)

type Service struct {
	debug      bool
	public     bool
	files      npncore.FileLoader
	user       user.Service
	auth       auth.Service
	logger     *logrus.Logger
	Session    *session.Service
	Collection *collection.Service
	Import     *imprt.Service
	Caller     *call.Service
	Socket     *npnconnection.Service
}

var _ npnweb.AppInfo = (*Service)(nil)

func NewService(debug bool, public bool, multiuser bool, files npncore.FileLoader, redir string, logger *logrus.Logger) *Service {
	us := userfs.NewServiceFilesystem(multiuser, files, logger)
	sessSvc := session.NewService(multiuser, files, logger)
	collSvc := collection.NewService(multiuser, files, logger)
	reqSvc := request.NewService(multiuser, files, logger)
	callSvc := call.NewService(sessSvc, logger)
	searchSvc := search.NewService(collSvc, reqSvc, logger)
	importSvc := imprt.NewService(files, logger)
	socketDeps := &socket.Dependencies{User: us, Session: sessSvc, Collection: collSvc, Request: reqSvc, Caller: callSvc, Search: searchSvc, Import: importSvc}

	return &Service{
		debug:      debug,
		public:     public,
		files:      files,
		user:       us,
		auth:       authfs.NewServiceFS(multiuser, redir, files, logger, us),
		logger:     logger,
		Session:    sessSvc,
		Collection: collSvc,
		Import:     importSvc,
		Caller:     callSvc,
		Socket:     socket.NewService(socketDeps, logger),
	}
}

func (s *Service) Debug() bool {
	return s.debug
}

func (s *Service) Files() npncore.FileLoader {
	return s.files
}

func (s *Service) User() user.Service {
	return s.user
}

func (s *Service) Auth() auth.Service {
	return s.auth
}

func (s *Service) Logger() *logrus.Logger {
	return s.logger
}

func (s *Service) Valid() bool {
	return true
}

func (s *Service) Public() bool {
	return s.public
}

func Svc(a npnweb.AppInfo) *Service {
	return a.(*Service)
}
