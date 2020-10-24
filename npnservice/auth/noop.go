package auth

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
)

type ServiceNoop struct{}

var _ Service = (*ServiceNoop)(nil)

func (s *ServiceNoop) FullURL(path string) string {
	return path
}

func NewServiceNoop() Service {
	svc := ServiceNoop{}
	return &svc
}

func (s *ServiceNoop) Enabled() bool {
	return false
}

func (s *ServiceNoop) EnabledProviders() Providers {
	return nil
}

func (s *ServiceNoop) MergeProfile(p *npnuser.UserProfile, record *Record) (*Record, error) {
	return record, nil
}

func (s *ServiceNoop) NewRecord(r *Record) (*Record, error) {
	panic("implement me")
}

func (s *ServiceNoop) UpdateRecord(r *Record) error {
	panic("implement me")
}

func (s *ServiceNoop) GetByProviderID(userID uuid.UUID, key string, code string) *Record {
	panic("implement me")
}

func (s *ServiceNoop) List(params *npncore.Params) Records {
	return nil
}

func (s *ServiceNoop) GetByID(userID uuid.UUID, authID uuid.UUID) *Record {
	return nil
}

func (s *ServiceNoop) GetByUserID(userID uuid.UUID, params *npncore.Params) Records {
	return nil
}

func (s *ServiceNoop) Delete(userID uuid.UUID, authID uuid.UUID) error {
	return errors.New("disabled auth service")
}
