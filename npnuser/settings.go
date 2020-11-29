package npnuser

type UserSettings struct {
	Mode  string `json:"mode"`
	NavB  string `json:"navB"`
	NavF  string `json:"navF"`
	MenuB string `json:"menuB"`
	MenuF string `json:"menuF"`
	MenuL string `json:"menuL"`
	BodyB string `json:"bodyB"`
	BodyL string `json:"bodyL"`
}

func (u *UserSettings) Clone() *UserSettings {
	return &UserSettings{
		Mode:  u.Mode,
		NavB:  u.NavB,
		NavF:  u.NavF,
		MenuB: u.MenuB,
		MenuF: u.MenuF,
		MenuL: u.MenuL,
		BodyB: u.BodyB,
		BodyL: u.BodyL,
	}
}

func (u *UserSettings) ModeCSS() string {
	if u == nil {
		return "uk-dark"
	}
	if u.Mode == "dark" {
		return "uk-light"
	}
	return "uk-dark"
}

func (u *UserSettings) Normalize() *UserSettings {
	if u == nil {
		return DefaultSettings
	}
	if u.Mode == "" {
		u.Mode = DefaultSettings.Mode
	}
	if u.NavB == "" {
		u.NavB = DefaultSettings.NavB
	}
	if u.NavF == "" {
		u.NavF = DefaultSettings.NavF
	}
	if u.MenuB == "" {
		u.MenuB = DefaultSettings.MenuB
	}
	if u.MenuF == "" {
		u.MenuF = DefaultSettings.MenuF
	}
	if u.MenuL == "" {
		u.MenuL = DefaultSettings.MenuL
	}
	if u.BodyB == "" {
		u.BodyB = DefaultSettings.BodyB
	}
	if u.BodyL == "" {
		u.BodyL = DefaultSettings.BodyL
	}
	return u
}

var DefaultSettings = &UserSettings{
	Mode:  "light",
	NavB:  "#193441",
	NavF:  "#dddddd",
	MenuB: "#3e606f",
	MenuF: "#ffffff",
	MenuL: "#cccccc",
	BodyB: "#fcfff5",
	BodyL: "#444",
}
