package auth

import (
	"encoding/json"
	"time"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

var slackScopes = []string{"users:read", "team:read"}

func slackAuth(tok *oauth2.Token) (*Record, error) {
	profile, err := loadProfile(tok)
	if err != nil {
		return nil, errors.Wrap(err, "error getting Slack user profile")
	}

	tm, err := loadTeam(tok)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Slack team")
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

func loadProfile(tok *oauth2.Token) (*slackProfile, error) {
	contents, err := callHTTP("https://slack.com/api/users.profile.get", tok.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "error reading Slack profile response")
	}

	var rsp = slackProfileResponse{}
	err = json.Unmarshal(contents, &rsp)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Slack profile")
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

func loadTeam(tok *oauth2.Token) (*slackTeam, error) {
	contents, err := callHTTP("https://slack.com/api/team.info", tok.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "error reading Slack team response")
	}

	var rsp = slackTeamResponse{}
	err = json.Unmarshal(contents, &rsp)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling Slack team")
	}

	return rsp.Team, nil
}
