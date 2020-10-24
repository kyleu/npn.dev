package authfs

import (
	"strings"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"logur.dev/logur"
)

type ServiceFS struct {
	enabled          bool
	enabledProviders auth.Providers
	redir            string
	files            npncore.FileLoader
	logger           logur.Logger
	users            user.Service
}

var _ auth.Service = (*ServiceFS)(nil)

func NewServiceFS(enabled bool, redir string, files npncore.FileLoader, logger logur.Logger, users user.Service) auth.Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeyAuth})

	if !strings.HasPrefix(redir, "http") {
		redir = "https://" + redir
	}
	if !strings.HasSuffix(redir, "/") {
		redir += "/"
	}

	svc := &ServiceFS{
		enabled: enabled,
		redir:   redir,
		files:   files,
		logger:  logger,
		users:   users,
	}

	for _, p := range auth.AllProviders {
		cfg := auth.GetConfig(svc, p)
		if cfg != nil {
			svc.enabledProviders = append(svc.enabledProviders, p)
		}
	}
	if len(svc.enabledProviders) == 0 {
		svc.enabled = false
	} else {
		logger.Info("auth service started for [" + strings.Join(svc.enabledProviders.Names(), ", ") + "]")
	}

	return svc
}

func (s *ServiceFS) Enabled() bool {
	return s.enabled
}

func (s *ServiceFS) EnabledProviders() auth.Providers {
	return s.enabledProviders
}

func (s *ServiceFS) FullURL(path string) string {
	return s.redir + strings.TrimPrefix(path, "/")
}
