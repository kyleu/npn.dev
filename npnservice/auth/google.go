package auth

import (
	"encoding/json"
	"time"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

var googleScopes = []string{
	"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/userinfo.email",
}

type googleUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func googleAuth(tok *oauth2.Token) (*Record, error) {
	contents, err := callHTTP("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken, "")
	if err != nil {
		return nil, errors.Wrap(err, "error reading Google response")
	}

	var user = googleUser{}
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Google user")
	}

	ret := Record{
		ID:         npncore.UUID(),
		Provider:   &ProviderGoogle,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
