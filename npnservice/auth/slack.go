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

var slackScopes = []string{"users:read", "team:read"}

func slackAuth(tok *oauth2.Token) (*Record, error) {
	client := &http.Client{}

	profile, err := loadProfile(tok, client)
	if err != nil {
		return nil, errors.Wrap(err, "error getting Slack user profile")
	}

	tm, err := loadTeam(tok, client)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling slack user")
	}

	ret := Record{
		ID:           npncore.UUID(),
		Provider:     &ProviderSlack,
		ProviderID:   profile.Email,
		UserListID:   tm.ID,
		UserListName: tm.Name,
		AccessToken:  tok.AccessToken,
		Expires:      &tok.Expiry,
		Name:         profile.Name,
		Email:        profile.Email,
		Picture:      profile.Picture,
		Created:      time.Time{},
	}
	return &ret, nil
}

type slackProfileResponse struct {
	Ok      bool          `json:"ok"`
	Profile *slackProfile `json:"profile"`
}

type slackProfile struct {
	Email   string `json:"email"`
	Name    string `json:"real_name"`
	Picture string `json:"image_192"`
}

func loadProfile(tok *oauth2.Token, client *http.Client) (*slackProfile, error) {
	req, err := http.NewRequest("GET", "https://slack.com/api/users.profile.get", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response from Slack")
	}

	var rsp = slackProfileResponse{}
	err = json.Unmarshal(contents, &rsp)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling slack user")
	}

	return rsp.Profile, nil
}

type slackTeamResponse struct {
	Ok   bool       `json:"ok"`
	Team *slackTeam `json:"team"`
}

type slackTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func loadTeam(tok *oauth2.Token, client *http.Client) (*slackTeam, error) {
	req, err := http.NewRequest("GET", "https://slack.com/api/team.info", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var rsp = slackTeamResponse{}
	err = json.Unmarshal(contents, &rsp)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling slack user")
	}

	return rsp.Team, nil
}
