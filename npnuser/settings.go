package npnuser

type UserSettings struct {
	Mode    string `json:"mode"`
	NavB    string `json:"navB"`
	NavF    string `json:"navF"`
	MenuB   string `json:"menuB"`
	MenuF   string `json:"menuF"`
	MenuL   string `json:"menuL"`
	BodyB   string `json:"bodyB"`
	BodyL   string `json:"bodyL"`
}

func (u *UserSettings) Clone() *UserSettings {
	return &UserSettings{
		Mode:    u.Mode,
		NavB:    u.NavB,
		NavF:    u.NavF,
		MenuB:   u.MenuB,
		MenuF:   u.MenuF,
		MenuL:   u.MenuL,
		BodyB:   u.BodyB,
		BodyL:   u.BodyL,
	}
}

func (u *UserSettings) ModeCSS() string {
	if u.Mode == "dark" {
		return "uk-light"
	}
	return "uk-dark"
}


var DefaultSettings = &UserSettings{
	Mode:    "light",
	NavB:    "#193441",
	NavF:    "#fff",
	MenuB:   "#3e606f",
	MenuF:   "#fff",
	MenuL:   "#ccc",
	BodyB:   "#fcfff5",
	BodyL:   "#444",
}
