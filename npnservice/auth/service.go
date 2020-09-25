package auth

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

type Service interface {
	Enabled() bool
	EnabledProviders() Providers

	URLFor(state string, prv *Provider) string
	FullURL(path string) string
	GetToken(prv *Provider, code string) (*oauth2.Token, error)
	Handle(profile *npnuser.UserProfile, prv *Provider, code string) (*Record, error)

	List(params *npncore.Params) Records
	GetByID(authID uuid.UUID) *Record
	GetByUserID(userID uuid.UUID, params *npncore.Params) Records
	GetDisplayByUserID(userID uuid.UUID, params *npncore.Params) (Records, Displays)
	Delete(authID uuid.UUID) error
}

func DecodeRecord(s Service, prv *Provider, code string) (*Record, error) {
	tok, err := s.GetToken(prv, code)
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
