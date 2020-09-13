package auth

import (
	"context"
	"github.com/kyleu/npn/npncore"
	"os"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

func (s *Service) callbackURL(k string) string {
	return s.FullURL("auth/callback/" + k)
}

func (s *Service) getConfig(prv *Provider) *oauth2.Config {
	idKey, secretKey := envsFor(prv)
	id := os.Getenv(idKey)
	secret := os.Getenv(secretKey)
	if len(id) == 0 || len(secret) == 0 {
		return nil
	}

	ret := oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		Endpoint:     prv.Endpoint,
		RedirectURL:  s.callbackURL(prv.Key),
		Scopes:       prv.Scopes,
	}

	return &ret
}

func (s *Service) URLFor(state string, prv *Provider) string {
	cfg := s.getConfig(prv)
	if cfg == nil {
		return ""
	}
	return cfg.AuthCodeURL(state)
}

func (s *Service) getToken(prv *Provider, code string) (*oauth2.Token, error) {
	cfg := s.getConfig(prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	ctx := context.TODO()
	return cfg.Exchange(ctx, code)
}

func (s *Service) decodeRecord(prv *Provider, code string) (*Record, error) {
	tok, err := s.getToken(prv, code)
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

func envsFor(prv *Provider) (string, string) {
	var id = npncore.AppKey + "_client_id_" + prv.Key
	var secret = npncore.AppKey + "_client_secret_" + prv.Key
	return id, secret
}
