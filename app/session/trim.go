package session

import (
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/header"
)

func (s *Session) Normalize(key string) *Session {
	if s == nil {
		return nil
	}
	if len(key) > 0 {
		s.Key = key
	}
	if s.Key == "" {
		s.Key = "untitled-" + npncore.RandomString(6)
	}

	s.Trim()

	return s
}

func (s *Session) Minify() *Session {
	s.Trim()

	if len(s.Cookies) == 0 {
		s.Cookies = nil
	}
	if len(s.Variables) == 0 {
		s.Variables = nil
	}

	return s
}

func (s *Session) Trim() {
	cookies := make(header.Cookies, 0, len(s.Cookies))
	for _, x := range s.Cookies {
		if len(x.Name) > 0 {
			cookies = append(cookies, x)
		}
	}
	s.Cookies = cookies

	variables := make(Variables, 0, len(s.Variables))
	for _, x := range s.Variables {
		if len(x.Key) > 0 {
			variables = append(variables, x)
		}
	}
	s.Variables = variables
}
