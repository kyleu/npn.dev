package npnuser

import (
	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

// Global ID used by the system UserProfile
var SystemUserID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000")

// Represent's a user and their settings
type UserProfile struct {
	UserID   uuid.UUID
	Name     string
	Role     Role
	Settings *UserSettings
	Picture  string
	Locale   language.Tag
}

// Constructor
func NewUserProfile(userID uuid.UUID, name string) *UserProfile {
	return &UserProfile{
		UserID:   userID,
		Name:     name,
		Role:     RoleGuest,
		Settings: DefaultSettings.Clone(),
		Locale:   language.AmericanEnglish,
	}
}

// Converts this UserProfile into a Profile
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

// Represent's a user and their settings, intended for JSON representation
type Profile struct {
	UserID   uuid.UUID     `json:"userID"`
	Name     string        `json:"name"`
	Role     string        `json:"role"`
	Settings *UserSettings `json:"settings"`
	Picture  string        `json:"picture"`
	Locale   string        `json:"locale,omitempty"`
}

// Converts this Profile into a UserProfile
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
