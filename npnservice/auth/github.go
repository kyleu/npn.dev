package auth

import (
	"encoding/json"
	"time"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

var githubScopes = []string{"profile"}

type githubUser struct {
	ID      string `json:"login"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"avatar_url"`
}

func githubAuth(tok *oauth2.Token) (*Record, error) {
	contents, err := callHTTP("https://api.github.com/user", tok.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "error reading Github response")
	}

	var user = githubUser{}
	err = json.Unmarshal(contents, &user)

	if err != nil {
		return nil, errors.Wrap(err, "error marshalling GitHub user")
	}

	ret := Record{
		ID:         npncore.UUID(),
		Provider:   &ProviderGitHub,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
