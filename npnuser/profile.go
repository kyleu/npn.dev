package npnuser

import (
	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

type UserProfile struct {
	UserID    uuid.UUID
	Name      string
	Theme     Theme
	Role      Role
	Settings  *UserSettings
	Picture   string
	Locale    language.Tag
}

func NewUserProfile(userID uuid.UUID, name string) *UserProfile {
	return &UserProfile{
		UserID:   userID,
		Name:     name,
		Theme:    ThemeAuto,
		Role:     RoleGuest,
		Settings: DefaultSettings.Clone(),
		Locale:   language.AmericanEnglish,
	}
}

type Profile struct {
	UserID   uuid.UUID     `json:"userID"`
	Name     string        `json:"name"`
	Theme    string        `json:"theme"`
	Role     string        `json:"role"`
	Settings *UserSettings `json:"settings"`
	Picture  string        `json:"picture"`
	Locale   string        `json:"locale,omitempty"`
}

func (p *UserProfile) ToProfile() *Profile {
	return &Profile{
		UserID:   p.UserID,
		Name:     p.Name,
		Theme:    p.Theme.String(),
		Role:     p.Role.String(),
		Settings: p.Settings,
		Picture:  p.Picture,
		Locale:   p.Locale.String(),
	}
}
