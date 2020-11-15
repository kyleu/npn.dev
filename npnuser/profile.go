package npnuser

import (
	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

type UserProfile struct {
	UserID   uuid.UUID
	Name     string
	Role     Role
	Settings *UserSettings
	Picture  string
	Locale   language.Tag
}

func NewUserProfile(userID uuid.UUID, name string) *UserProfile {
	return &UserProfile{
		UserID:   userID,
		Name:     name,
		Role:     RoleGuest,
		Settings: DefaultSettings.Clone(),
		Locale:   language.AmericanEnglish,
	}
}

func (p *UserProfile) ToProfile() *Profile {
	return &Profile{
		UserID:   p.UserID,
		Name:     p.Name,
		Role:     p.Role.String(),
		Settings: p.Settings,
		Picture:  p.Picture,
		Locale:   p.Locale.String(),
	}
}

type Profile struct {
	UserID   uuid.UUID     `json:"userID"`
	Name     string        `json:"name"`
	Role     string        `json:"role"`
	Settings *UserSettings `json:"settings"`
	Picture  string        `json:"picture"`
	Locale   string        `json:"locale,omitempty"`
}

func (p *Profile) ToUserProfile() *UserProfile {
	loc, err := language.Parse(p.Locale)
	if err != nil {
		loc = language.AmericanEnglish
	}
	return &UserProfile{
		UserID:   p.UserID,
		Name:     p.Name,
		Role:     RoleFromString(p.Role),
		Settings: p.Settings,
		Picture:  p.Picture,
		Locale:   loc,
	}
}
