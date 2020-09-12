package auth

const KeyBasic = "basic"

type Auth interface {
	Type() string
}

type Auths []Auth

func (a Auths) HasBasic() bool {
	for _, x := range a {
		if x.Type() == KeyBasic {
			return true
		}
	}
	return false
}

func (a Auths) GetBasic() (string, string) {
	for _, x := range a {
		if x.Type() == KeyBasic {
			b := x.(*Basic)
			return b.Username, b.Password
		}
	}
	return "", ""
}

type Basic struct {
	Username string
	Password string
}

func (*Basic) Type() string {
	return KeyBasic
}
