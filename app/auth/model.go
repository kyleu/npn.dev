package auth

import (
	"encoding/json"
	"fmt"

	"github.com/kyleu/libnpn/npncore"
	"logur.dev/logur"

	"emperror.dev/errors"
)

type Config interface {
	GetType() string
	String() string
	Clone() *Auth
}

type Auth struct {
	Type   string `json:"type"`
	Config Config `json:"config"`
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

func (a *Auth) UnmarshalJSON(data []byte) error {
	x := &authJSON{}
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	a.Type = x.Type
	switch a.Type {
	case KeyBasic:
		basic := &Basic{}
		err = json.Unmarshal(x.Config, &basic)
		if err != nil {
			return err
		}
		a.Config = basic
	default:
		return errors.New("invalid auth type [" + x.Type + "]")
	}
	return nil
}

func (a *Auth) GetBasic() (string, string) {
	if a != nil && a.IsBasic() {
		b := a.Config.(*Basic)
		return b.Username, b.Password
	}
	return "", ""
}

func (a *Auth) IsBasic() bool {
	return a != nil && a.Type == KeyBasic
}

func (a *Auth) Merge(data npncore.Data, logger logur.Logger) *Auth {
	if a == nil {
		return nil
	}
	return &Auth{
		Type:   npncore.MergeLog("auth.type", a.Type, data, logger),
		Config: a.Config, // TODO
	}
}

func (a *Auth) Clone() *Auth {
	if a == nil || a.Config == nil {
		return nil
	}
	return a.Config.Clone()
}
