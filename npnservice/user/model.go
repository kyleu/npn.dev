package user

import (
	"time"

	"github.com/kyleu/npn/npnuser"

	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

type SystemUser struct {
	UserID    uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Role      string    `db:"role"`
	Theme     string    `db:"theme"`
	NavColor  string    `db:"nav_color"`
	LinkColor string    `db:"link_color"`
	Picture   string    `db:"picture"`
	Locale    string    `db:"locale"`
	Created   time.Time `db:"created"`
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
		Theme:     npnuser.ThemeFromString(su.Theme),
		NavColor:  su.NavColor,
		LinkColor: su.LinkColor,
		Picture:   su.Picture,
		Locale:    locale,
	}
}
