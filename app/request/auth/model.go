package auth

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
)

type AuthConfig interface {
	GetType() string
	String() string
}

type Auth struct {
	Type   string     `json:"type"`
	Config AuthConfig `json:"config"`
}

func (a *Auth) String() string {
	s := "nil"
	if a.Config != nil {
		s = a.Config.String()
	}
	return fmt.Sprintf("[%v]%v", a.Type, s)
}

type authJSON struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

func (w *Auth) UnmarshalJSON(data []byte) error {
	x := &authJSON{}
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	w.Type = x.Type
	switch w.Type {
	case KeyBasic:
		basic := &Basic{}
		err = json.Unmarshal(x.Config, &basic)
		if err != nil {
			return err
		}
		w.Config = basic
	default:
		return errors.New("invalid auth type [" + x.Type + "]")
	}
	return nil
}

type Auths []*Auth

func (a Auths) HasBasic() bool {
	for _, x := range a {
		if x.Type == KeyBasic {
			return true
		}
	}
	return false
}

func (a Auths) GetBasic() (string, string) {
	for _, x := range a {
		if x.Type == KeyBasic {
			b := x.Config.(*Basic)
			return b.Username, b.Password
		}
	}
	return "", ""
}
