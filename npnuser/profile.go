package npnuser

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

type Role struct {
	Key string
}

var RoleGuest = Role{
	Key: "guest",
}

var RoleUser = Role{
	Key: "user",
}

var RoleAdmin = Role{
	Key: "admin",
}

var AllRoles = []Role{RoleGuest, RoleUser, RoleAdmin}

func RoleFromString(s string) Role {
	for _, t := range AllRoles {
		if t.Key == s {
			return t
		}
	}
	return RoleGuest
}

func (t *Role) String() string {
	return t.Key
}

func (t *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Role) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = RoleFromString(s)
	return nil
}

var (
	DefaultNavColor  = "bluegrey"
	DefaultLinkColor = "bluegrey"
)

type UserProfile struct {
	UserID    uuid.UUID
	Name      string
	Theme     Theme
	Role      Role
	NavColor  string
	LinkColor string
	Picture   string
	Locale    language.Tag
}

func NewUserProfile(userID uuid.UUID, name string) *UserProfile {
	return &UserProfile{
		UserID:    userID,
		Name:      name,
		Theme:     ThemeAuto,
		Role:      RoleGuest,
		NavColor:  DefaultNavColor,
		LinkColor: DefaultLinkColor,
		Locale:    language.AmericanEnglish,
	}
}

type Profile struct {
	UserID    uuid.UUID `json:"userID"`
	Name      string    `json:"name"`
	Theme     string    `json:"theme"`
	Role      string    `json:"role"`
	NavColor  string    `json:"navColor"`
	LinkColor string    `json:"linkColor"`
	Picture   string    `json:"picture"`
	Locale    string    `json:"locale,omitempty"`
}

func (p *UserProfile) ToProfile() *Profile {
	return &Profile{
		UserID:    p.UserID,
		Name:      p.Name,
		Theme:     p.Theme.String(),
		Role:      p.Role.String(),
		NavColor:  p.NavColor,
		LinkColor: p.LinkColor,
		Picture:   p.Picture,
		Locale:    p.Locale.String(),
	}
}
