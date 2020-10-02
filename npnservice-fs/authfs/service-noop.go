package authfs

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnuser"
	"golang.org/x/oauth2"
)

type ServiceNoop struct{}

var _ auth.Service = (*ServiceNoop)(nil)

func (s *ServiceNoop) FullURL(path string) string {
	return path
}

func NewServiceNoop() auth.Service {
	svc := ServiceNoop{}
	return &svc
}

func (s *ServiceNoop) Enabled() bool {
	return false
}

func (s *ServiceNoop) EnabledProviders() auth.Providers {
	return nil
}

func (s *ServiceNoop) URLFor(state string, prv *auth.Provider) string {
	return "invalid"
}

func (s *ServiceNoop) GetToken(prv *auth.Provider, code string) (*oauth2.Token, error) {
	return nil, nil
}

func (s *ServiceNoop) Handle(profile *npnuser.UserProfile, prv *auth.Provider, code string) (*auth.Record, error) {
	return nil, nil
}

func (s *ServiceNoop) List(params *npncore.Params) auth.Records {
	return nil
}

func (s *ServiceNoop) GetByID(authID uuid.UUID) *auth.Record {
	return nil
}

func (s *ServiceNoop) GetByUserID(userID uuid.UUID, params *npncore.Params) auth.Records {
	return nil
}

func (s *ServiceNoop) GetDisplayByUserID(userID uuid.UUID, params *npncore.Params) (auth.Records, auth.Displays) {
	return nil, nil
}

func (s *ServiceNoop) Delete(authID uuid.UUID) error {
	return errors.New("disabled auth service")
}
