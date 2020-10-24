package auth

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"golang.org/x/oauth2"
)

func CallbackURL(s Service, k string) string {
	return s.FullURL("auth/callback/" + k)
}

func GetConfig(s Service, prv *Provider) *oauth2.Config {
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
		RedirectURL:  CallbackURL(s, prv.Key),
		Scopes:       prv.Scopes,
	}

	return &ret
}

func URLFor(s Service, state string, prv *Provider) string {
	cfg := GetConfig(s, prv)
	if cfg == nil {
		return ""
	}
	return cfg.AuthCodeURL(state)
}

func GetDisplayByUserID(s Service, userID uuid.UUID, params *npncore.Params) (Records, Displays) {
	if !s.Enabled() {
		return nil, nil
	}

	rec := s.GetByUserID(userID, params)
	disp := make(Displays, 0, len(rec))
	for _, r := range rec {
		disp = append(disp, r.ToDisplay())
	}
	return rec, disp
}

func getToken(s Service, prv *Provider, code string) (*oauth2.Token, error) {
	cfg := GetConfig(s, prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	ctx := context.Background()
	return cfg.Exchange(ctx, code)
}

func envsFor(prv *Provider) (string, string) {
	var id = npncore.AppKey + "_client_id_" + prv.Key
	var secret = npncore.AppKey + "_client_secret_" + prv.Key
	return id, secret
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
