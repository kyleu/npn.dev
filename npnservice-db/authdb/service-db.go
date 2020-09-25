package authdb

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnservice/auth"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnservice/user"

	"logur.dev/logur"
)

type ServiceDatabase struct {
	enabled          bool
	enabledProviders auth.Providers
	redir            string
	db               *npndatabase.Service
	logger           logur.Logger
	users            user.Service
}

func NewServiceDatabase(enabled bool, redir string /* actions *action.Service, */, db *npndatabase.Service, logger logur.Logger, users user.Service) auth.Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeyAuth})

	if !strings.HasPrefix(redir, "http") {
		redir = "https://" + redir
	}
	if !strings.HasSuffix(redir, "/") {
		redir += "/"
	}

	svc := ServiceDatabase{
		enabled: enabled,
		redir:   redir,
		db:      db,
		logger:  logger,
		users:   users,
	}

	for _, p := range auth.AllProviders {
		cfg := svc.getConfig(p)
		if cfg != nil {
			svc.enabledProviders = append(svc.enabledProviders, p)
		}
	}
	if len(svc.enabledProviders) == 0 {
		svc.enabled = false
	} else {
		logger.Info("auth service started for [" + strings.Join(svc.enabledProviders.Names(), ", ") + "]")
	}

	return &svc
}

func (s *ServiceDatabase) Enabled() bool {
	return s.enabled
}

func (s *ServiceDatabase) EnabledProviders() auth.Providers {
	return s.enabledProviders
}

func callHTTP(url string, auth string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(auth) > 0 {
		req.Header.Add("Authorization", "Bearer "+auth)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	return ioutil.ReadAll(response.Body)
}

func (s *ServiceDatabase) GetDisplayByUserID(userID uuid.UUID, params *npncore.Params) (auth.Records, auth.Displays) {
	if !s.Enabled() {
		return nil, nil
	}

	rec := s.GetByUserID(userID, params)
	disp := make(auth.Displays, 0, len(rec))
	for _, r := range rec {
		disp = append(disp, r.ToDisplay())
	}
	return rec, disp
}

func (s *ServiceDatabase) FullURL(path string) string {
	return s.redir + strings.TrimPrefix(path, "/")
}

