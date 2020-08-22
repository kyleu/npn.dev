package auth

import (
	"encoding/json"
	"time"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

var amazonScopes = []string{"profile"}

type amazonUser struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func amazonAuth(tok *oauth2.Token) (*Record, error) {
	contents, err := callHTTP("https://api.amazon.com/user/profile", tok.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "error reading Amazon response")
	}

	var user = amazonUser{}
	err = json.Unmarshal(contents, &user)

	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Amazon user")
	}

	ret := Record{
		ID:         npncore.UUID(),
		Provider:   &ProviderAmazon,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    "",
		Created:    time.Time{},
	}
	return &ret, nil
}
