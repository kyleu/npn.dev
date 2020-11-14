package user

import (
	"time"

	"github.com/kyleu/npn/npnuser"

	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

type SystemUser struct {
	UserID    uuid.UUID             `db:"id"`
	Name      string                `db:"name"`
	Role      string                `db:"role"`
	Settings  *npnuser.UserSettings `db:"settings"`
	Picture   string                `db:"picture"`
	Locale    string                `db:"locale"`
	Created   time.Time             `db:"created"`
}

type SystemUsers = []*SystemUser

func (su *SystemUser) ToProfile() *npnuser.UserProfile {
	locale, err := language.Parse(su.Locale)
	if err != nil {
		locale = language.AmericanEnglish
	}

	return &npnuser.UserProfile{
		UserID:    su.UserID,
		Name:      su.Name,
		Role:      npnuser.RoleFromString(su.Role),
		Settings:  su.Settings,
		Picture:   su.Picture,
		Locale:    locale,
	}
}

func FromProfile(p *npnuser.UserProfile, created time.Time) *SystemUser {
	return &SystemUser{
		UserID:    p.UserID,
		Name:      p.Name,
		Role:      p.Role.String(),
		Settings:  p.Settings,
		Picture:   p.Picture,
		Locale:    p.Locale.String(),
		Created:   created,
	}
}
