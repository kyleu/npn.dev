package auth

import (
	"fmt"
	"strings"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnservice/user"

	"github.com/gofrs/uuid"
	"logur.dev/logur"
)

type Service struct {
	Enabled          bool
	EnabledProviders Providers
	redir            string
	// actions          *action.Service
	db     *npndatabase.Service
	logger logur.Logger
	users  *user.Service
}

func NewService(enabled bool, redir string /* actions *action.Service, */, db *npndatabase.Service, logger logur.Logger, users *user.Service) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeyAuth})

	if !strings.HasPrefix(redir, "http") {
		redir = "https://" + redir
	}
	if !strings.HasSuffix(redir, "/") {
		redir += "/"
	}

	svc := Service{
		Enabled: enabled,
		redir:   redir,
		// actions: actions,
		db:     db,
		logger: logger,
		users:  users,
	}

	for _, p := range AllProviders {
		cfg := svc.getConfig(p)
		if cfg != nil {
			svc.EnabledProviders = append(svc.EnabledProviders, p)
		}
	}
	if len(svc.EnabledProviders) == 0 {
		svc.Enabled = false
	} else {
		logger.Info("auth service started for [" + strings.Join(svc.EnabledProviders.Names(), ", ") + "]")
	}

	return &svc
}

func (s *Service) GetDisplayByUserID(userID uuid.UUID, params *npncore.Params) (Records, Displays) {
	if !s.Enabled {
		return nil, nil
	}

	params = npncore.ParamsWithDefaultOrdering(npncore.KeyAuth, params, npncore.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := npndatabase.SQLSelect("*", npncore.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving auth entries for user [%v]: %+v", userID, err))
		return nil, nil
	}
	rec := make(Records, 0, len(dtos))
	for _, dto := range dtos {
		rec = append(rec, dto.toRecord())
	}
	disp := make(Displays, 0, len(rec))
	for _, r := range rec {
		disp = append(disp, r.ToDisplay())
	}
	return rec, disp
}

func (s *Service) FullURL(path string) string {
	return s.redir + strings.TrimPrefix(path, "/")
}
