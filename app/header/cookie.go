package header

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
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
	HTTPOnly bool      `json:"httpOnly,omitempty"`
	SameSite string    `json:"sameSite,omitempty"`
}

func NewCookie(c *http.Cookie) *Cookie {
	ss := ""
	switch c.SameSite {
	case http.SameSiteLaxMode:
		ss = "lax"
	case http.SameSiteStrictMode:
		ss = "strict"
	case http.SameSiteNoneMode:
		ss = "none"
	case http.SameSiteDefaultMode:
		ss = "default"
	}
	return &Cookie{
		Name:     c.Name,
		Value:    c.Value,
		Path:     c.Path,
		Domain:   c.Domain,
		Expires:  c.Expires,
		MaxAge:   c.MaxAge,
		Secure:   c.Secure,
		HTTPOnly: c.HttpOnly,
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
	return c.NativeSmall().String()
}

func (c *Cookie) NativeSmall() *http.Cookie {
	ret := &http.Cookie{Name: c.Name, Value: c.Value}
	return ret
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
	ret := &http.Cookie{
		Name:       c.Name,
		Value:      c.Value,
		Path:       c.Path,
		Domain:     c.Domain,
		Expires:    c.Expires,
		RawExpires: c.Expires.Format(time.RFC1123),
		MaxAge:     c.MaxAge,
		Secure:     c.Secure,
		HttpOnly:   c.HTTPOnly,
		SameSite:   ss,
	}
	return ret
}

func (c *Cookie) Matches(x *Cookie) bool {
	return c.Name == x.Name
}

func (c *Cookie) Equals(x *Cookie) bool {
	if c == nil || x == nil {
		return false
	}
	return *c == *x
}

type Cookies []*Cookie

func (c Cookies) String() string {
	ret := make([]string, 0, len(c))
	for _, cook := range c {
		ret = append(ret, cook.String())
	}
	return strings.Join(ret, "; ")
}

func (c Cookies) Native() []*http.Cookie {
	ret := make([]*http.Cookie, 0, len(c))
	for _, cook := range c {
		ret = append(ret, cook.Native())
	}
	return ret
}

func (c Cookies) Qualifying(u *url.URL) Cookies {
	x, _ := cookiejar.New(nil)
	x.SetCookies(u, c.Native())

	q := x.Cookies(u)

	ret := make(Cookies, 0, len(q))
	for _, cook := range q {
		ret = append(ret, NewCookie(cook))
	}
	return ret
}
