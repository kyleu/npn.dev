package session

import (
	"github.com/kyleu/npn/app/header"
)

type Session struct {
	Key       string     `json:"key"`
	Cookies   header.Cookies   `json:"cookies"`
	Variables Variables `json:"variables"`
}

func (s Session) AddCookies(n header.Cookies) {
	println("COOOOOOOOOOOKIES!")
	println(n)
}
