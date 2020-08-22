package auth

import (
	"encoding/json"
	"time"

	"github.com/kyleu/npn/npncore"

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
	contents, err := callHTTP("https://graph.microsoft.com/v1.0/me/", tok.AccessToken)
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
