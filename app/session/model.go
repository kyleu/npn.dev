package session

import (
	"github.com/kyleu/npn/app/header"
)

type Summary struct {
	Key           string `json:"key"`
	Title         string `json:"title,omitempty"`
	CookieCount   int    `json:"cookieCount"`
	VariableCount int    `json:"variableCount"`
}

type Summaries []*Summary

type Session struct {
	Key       string         `json:"key"`
	Title     string         `json:"title,omitempty"`
	Cookies   header.Cookies `json:"cookies"`
	Variables Variables      `json:"variables"`
}

type Sessions []*Session

func (s Session) AddCookies(n header.Cookies) {
	println("COOOOOOOOOOOKIES!")
	println(n)
}

func (s Session) ToSummary() *Summary {
	return &Summary{
		Key:           s.Key,
		Title:         s.Title,
		CookieCount:   len(s.Cookies),
		VariableCount: len(s.Variables),
	}
}
