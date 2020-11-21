package header

import (
	"fmt"
	"net/http"
	"time"
)

type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path,omitempty"`
	Domain   string    `json:"domain,omitempty"`
	Expires  time.Time `json:"expires,omitempty"`
	MaxAge   int       `json:"maxAge,omitempty"`
	Secure   bool      `json:"secure,omitempty"`
	HttpOnly bool      `json:"httpOnly,omitempty"`
	SameSite string    `json:"sameSite,omitempty"`
}

func NewCookie(c *http.Cookie) *Cookie {
	ss := "default"
	switch c.SameSite {
	case http.SameSiteLaxMode:
		ss = "lax"
	case http.SameSiteStrictMode:
		ss = "strict"
	case http.SameSiteNoneMode:
		ss = "none"
	}
	return &Cookie{
		Name:     c.Name,
		Value:    c.Value,
		Path:     c.Path,
		Domain:   c.Domain,
		Expires:  c.Expires,
		MaxAge:   c.MaxAge,
		Secure:   c.Secure,
		HttpOnly: c.HttpOnly,
		SameSite: ss,
	}
}

func ParseCookies(cs []*http.Cookie) Cookies {
	ret := Cookies{}
	for _, c := range cs {
		ret = append(ret, NewCookie(c))
	}
	return ret
}

func (c *Cookie) String() string {
	return fmt.Sprintf("{%v}", c.Name)
}

func (c *Cookie) Native() *http.Cookie {
	ss := http.SameSiteDefaultMode
	switch c.SameSite {
	case "lax":
		ss = http.SameSiteLaxMode
	case "strict":
		ss = http.SameSiteStrictMode
	case "none":
		ss = http.SameSiteNoneMode
	}
	return &http.Cookie{
		Name:       c.Name,
		Value:      c.Value,
		Path:       c.Path,
		Domain:     c.Domain,
		Expires:    c.Expires,
		RawExpires: c.Expires.Format(time.RFC1123),
		MaxAge:     c.MaxAge,
		Secure:     c.Secure,
		HttpOnly:   c.HttpOnly,
		SameSite:   ss,
		Raw:        c.String(),
	}
}

type Cookies []*Cookie
