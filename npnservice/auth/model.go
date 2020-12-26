package auth

import (
	"time"

	"golang.org/x/oauth2/facebook"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
	"golang.org/x/oauth2/slack"

	"emperror.dev/errors"

	"github.com/gofrs/uuid"
)

// An auth provider, such as Google or Facebook
type Provider struct {
	Key      string          `json:"key"`
	Title    string          `json:"title,omitempty"`
	Icon     string          `json:"icon,omitempty"`
	Endpoint oauth2.Endpoint `json:"-"`
	Scopes   []string        `json:"scopes,omitempty"`
}

// Array helper
type Providers []*Provider

// Returns an array of names for these Providers
func (p Providers) Names() []string {
	ret := make([]string, 0, len(p))
	for _, prv := range p {
		ret = append(ret, prv.Title)
	}
	return ret
}

var ProviderGitHub = Provider{Key: "github", Title: "GitHub", Icon: "github-alt", Endpoint: github.Endpoint, Scopes: githubScopes}
var ProviderGoogle = Provider{Key: "google", Title: "Google", Icon: "google", Endpoint: google.Endpoint, Scopes: googleScopes}
var ProviderSlack = Provider{Key: "slack", Title: "Slack", Icon: "hashtag", Endpoint: slack.Endpoint, Scopes: slackScopes}
var ProviderFacebook = Provider{Key: "facebook", Title: "Facebook", Icon: "facebook", Endpoint: facebook.Endpoint, Scopes: facebookScopes}
var ProviderAmazon = Provider{Key: "amazon", Title: "Amazon", Icon: "cart", Endpoint: amazon.Endpoint, Scopes: amazonScopes}
var ProviderMicrosoft = Provider{Key: "microsoft", Title: "Microsoft", Icon: "world", Endpoint: microsoft.AzureADEndpoint(""), Scopes: microsoftScopes}

var AllProviders = Providers{&ProviderGitHub, &ProviderGoogle, &ProviderSlack, &ProviderFacebook, &ProviderAmazon, &ProviderMicrosoft}

// Finds the Provider matching the provided Key
func ProviderFromString(s string) *Provider {
	for _, t := range AllProviders {
		if t.Key == s {
			return t
		}
	}
	return &ProviderGitHub
}

// A JSON display representation of a Record
type Display struct {
	ID            uuid.UUID `json:"id"`
	Provider      string    `json:"provider"`
	Email         string    `json:"email"`
	ProvidesUsers string    `json:"providesUsers"`
}

// Array helper
type Displays []*Display

// An auth Record instance from a Provider
type Record struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Provider     *Provider
	ProviderID   string
	UserListID   string
	UserListName string
	AccessToken  string
	Expires      *time.Time
	Name         string
	Email        string
	Picture      string
	Created      time.Time
}

// Array helper
type Records []*Record

// Converts this Record to a Display, hiding sensitive information
func (r *Record) ToDisplay() *Display {
	return &Display{
		ID:            r.ID,
		Provider:      r.Provider.Key,
		Email:         r.Email,
		ProvidesUsers: r.UserListName,
	}
}

// Filters these Records to those matching the provided key
func (r Records) FindByProvider(key string) Records {
	var ret Records
	for _, e := range r {
		if e.Provider.Key == key {
			ret = append(ret, e)
		}
	}
	return ret
}

// Returns an array of email addresses for these Records
func (r Records) Emails() []string {
	ret := make([]string, 0, len(r))
	for _, e := range r {
		ret = append(ret, e.Email)
	}
	return ret
}

// Error used with auth isn't enabled
var ErrorAuthDisabled = errors.New("authorization is disabled")
