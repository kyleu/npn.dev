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

var defaultSession = &Session{
	Key:       "_",
	Title:     "Default Session",
	Cookies:   make(header.Cookies, 0),
	Variables: make(Variables, 0),
}

type Sessions []*Session

func (s *Session) AddCookies(c ...*header.Cookie) bool {
	modified := false

	for _, in := range c {
		matched := false
		for xIdx, x := range s.Cookies {
			if in.Matches(x) {
				matched = true
				if !in.Equals(x) {
					modified = true
				}
				s.Cookies[xIdx] = in
			}
		}
		if !matched {
			s.Cookies = append(s.Cookies, in)
			modified = true
		}
	}

	return modified
}

func (s *Session) AddVariables(v ...*Variable) bool {
	modified := false

	for _, in := range v {
		matched := false
		for xIdx, x := range s.Variables {
			if in.Matches(x) {
				matched = true
				if !in.Equals(x) {
					modified = true
				}
				s.Variables[xIdx] = in
			}
		}
		if !matched {
			s.Variables = append(s.Variables, in)
			modified = true
		}
	}

	return modified
}

func (s *Session) ToSummary() *Summary {
	return &Summary{
		Key:           s.Key,
		Title:         s.Title,
		CookieCount:   len(s.Cookies),
		VariableCount: len(s.Variables),
	}
}
