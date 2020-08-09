package auth

import (
	"encoding/json"
	"github.com/kyleu/npn/npncore"
	"io/ioutil"
	"net/http"
	"time"

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
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)

	// req.Header.Add("Authorization", "token "+tok.AccessToken)
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading GitHub response")
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
