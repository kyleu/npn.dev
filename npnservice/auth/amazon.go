package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.amazon.com/user/profile", nil)
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
