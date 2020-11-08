package npnuser

var (
	DefaultNavColor  = "bluegrey"
	DefaultLinkColor = "bluegrey"
)

type UserSettings struct {
	NavColor  string `json:"navColor"`
	LinkColor string `json:"linkColor"`
}

func (u *UserSettings) Clone() *UserSettings {
	return &UserSettings{NavColor: u.NavColor, LinkColor: u.LinkColor}
}

func (u *UserSettings) String() string {
	ret := "navColor => " + u.NavColor + ", linkColor => " + u.LinkColor
	return ret
}

var DefaultSettings = &UserSettings{
	NavColor:  DefaultNavColor,
	LinkColor: DefaultLinkColor,
}
