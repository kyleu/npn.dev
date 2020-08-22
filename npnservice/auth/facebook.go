package auth

import (
	"encoding/json"
	"time"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

var facebookScopes = []string{"public_profile", "email"}

type facebookData struct {
	URL string `json:"url"`
}

type facebookPicture struct {
	Data facebookData `json:"data"`
}

type facebookUser struct {
	ID      string          `json:"id"`
	Email   string          `json:"email"`
	Name    string          `json:"name"`
	Picture facebookPicture `json:"picture"`
}

func facebookAuth(tok *oauth2.Token) (*Record, error) {
	contents, err := callHTTP("https://graph.facebook.com/me?fields=id,name,email,picture&access_token="+tok.AccessToken, "")
	if err != nil {
		return nil, errors.Wrap(err, "error reading Facebook response")
	}

	var user = facebookUser{}
	err = json.Unmarshal(contents, &user)

	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Facebook user")
	}

	ret := Record{
		ID:         npncore.UUID(),
		Provider:   &ProviderFacebook,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture.Data.URL,
		Created:    time.Time{},
	}
	return &ret, nil
}
