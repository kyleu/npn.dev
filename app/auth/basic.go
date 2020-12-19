package auth

import (
	"fmt"
)

const KeyBasic = "basic"

type Basic struct {
	Username     string `json:"username"`
	Password     string `json:"password,omitempty"`
	ShowPassword bool   `json:"showPassword,omitempty"`
}

var _ Config = (*Basic)(nil)

func NewBasic(user string, pass string, showPass bool) *Auth {
	b := &Basic{Username: user, Password: pass, ShowPassword: showPass}
	return &Auth{Type: KeyBasic, Config: b}
}

func (b *Basic) GetType() string {
	return KeyBasic
}

func (b *Basic) String() string {
	passStr := ""
	if b.ShowPassword && len(b.Password) > 0 {
		passStr = ":" + b.Password
	}
	return fmt.Sprintf("%v%v", b.Username, passStr)
}

func (b *Basic) Clone() *Auth {
	return NewBasic(b.Username, b.Password, b.ShowPassword)
}
