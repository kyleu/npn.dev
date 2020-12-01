package auth

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
)

type Service interface {
	Enabled() bool
	EnabledProviders() Providers

	FullURL(path string) string
	MergeProfile(p *npnuser.UserProfile, record *Record) (*Record, error)

	NewRecord(r *Record) (*Record, error)
	UpdateRecord(r *Record) error
	Delete(userID uuid.UUID, authID uuid.UUID) error

	GetByID(userID uuid.UUID, authID uuid.UUID) *Record
	GetByProviderID(userID uuid.UUID, key string, code string) *Record
	GetByUserID(userID uuid.UUID, params *npncore.Params) Records
}

func DecodeRecord(s Service, prv *Provider, code string) (*Record, error) {
	tok, err := getToken(s, prv, code)
	if err != nil {
		return nil, errors.Wrap(err, "error getting token")
	}

	switch prv {
	case &ProviderGoogle:
		return googleAuth(tok)
	case &ProviderGitHub:
		return githubAuth(tok)
	case &ProviderSlack:
		return slackAuth(tok)
	case &ProviderFacebook:
		return facebookAuth(tok)
	case &ProviderAmazon:
		return amazonAuth(tok)
	case &ProviderMicrosoft:
		return microsoftAuth(tok)
	default:
		return nil, nil
	}
}
