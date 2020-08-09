package app

import (
	"github.com/kyleu/npn/app/parser"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
)

type Service struct {
	debug    bool
	Parsers  *parser.Parsers
	Schemata *schema.Service
	Projects *project.Service
	files    *npncore.FileLoader
	user     *user.Service
	auth     *auth.Service
	version  string
	commit   string
	logger   logur.Logger
}

func NewService(debug bool, version string, commitHash string, logger logur.Logger) *Service {
	files := npncore.NewFileLoader("./." + npncore.AppName, logger)
	us := user.NewService(files, nil, logger)
	au := auth.NewService(false, "", nil, logger, us)
	return &Service{
		debug:    debug,
		Parsers:  parser.NewParsers(logger),
		Schemata: schema.NewService(files, logger),
		Projects: project.NewService(files, logger),
		files:    files,
		user:     us,
		auth:     au,
		version:  version,
		commit:   commitHash,
		logger:   logger,
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

func Parsers(a npnweb.AppInfo) *parser.Parsers {
	return a.(*Service).Parsers
}

func Schemata(a npnweb.AppInfo) *schema.Service {
	return a.(*Service).Schemata
}

func Projects(a npnweb.AppInfo) *project.Service {
	return a.(*Service).Projects
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
