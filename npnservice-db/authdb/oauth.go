package authdb

import (
	"context"
	"os"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

func (s *ServiceDatabase) callbackURL(k string) string {
	return s.FullURL("auth/callback/" + k)
}

func (s *ServiceDatabase) getConfig(prv *auth.Provider) *oauth2.Config {
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

func (s *ServiceDatabase) URLFor(state string, prv *auth.Provider) string {
	cfg := s.getConfig(prv)
	if cfg == nil {
		return ""
	}
	return cfg.AuthCodeURL(state)
}

func (s *ServiceDatabase) GetToken(prv *auth.Provider, code string) (*oauth2.Token, error) {
	cfg := s.getConfig(prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	ctx := context.Background()
	return cfg.Exchange(ctx, code)
}

func envsFor(prv *auth.Provider) (string, string) {
	var id = npncore.AppKey + "_client_id_" + prv.Key
	var secret = npncore.AppKey + "_client_secret_" + prv.Key
	return id, secret
}
