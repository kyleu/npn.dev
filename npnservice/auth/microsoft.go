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

var microsoftScopes = []string{"user.read"}

type microsoftUser struct {
	ID      string `json:"id"`
	Email   string `json:"userPrincipalName"`
	Name    string `json:"displayName"`
	Picture string `json:"unused"`
}

func microsoftAuth(tok *oauth2.Token) (*Record, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading Microsoft response")
	}

	var user = microsoftUser{}
	err = json.Unmarshal(contents, &user)

	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Microsoft user")
	}

	ret := Record{
		ID:         npncore.UUID(),
		Provider:   &ProviderMicrosoft,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
