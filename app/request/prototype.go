package request

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/npncore"
)

type Prototype struct {
	Method   Method      `json:"method"`
	Protocol string      `json:"protocol"`
	Domain   string      `json:"domain"`
	Path     string      `json:"path,omitempty"`
	Query    QueryParams `json:"query,omitempty"`
	Fragment string      `json:"fragment,omitempty"`
	Headers  Headers     `json:"headers,omitempty"`
	Auth     auth.Auths  `json:"auth,omitempty"`
	Body     Body        `json:"body,omitempty"`
	Options  *Options    `json:"options,omitempty"`
}

func (p *Prototype) ToHTTP() *http.Request {
	p.URL()
	return &http.Request{
		Method:           p.Method.Key,
		URL:              p.URL(),
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             p.Domain,
		Form:             nil,
		PostForm:         nil,
		MultipartForm:    nil,
		Trailer:          nil,
		RemoteAddr:       "",
		RequestURI:       "",
		TLS:              nil,
		Cancel:           nil,
		Response:         nil,
	}
}

func PrototypeFromURL(method string, u *url.URL) *Prototype {
	var auths auth.Auths
	if u.User != nil {
		p, _ := u.User.Password()
		a := &auth.Basic{Username: u.User.Username(), Password: p}
		auths = auth.Auths{a}
	}
	return &Prototype{
		Method:   MethodGet,
		Protocol: u.Scheme,
		Domain:   u.Host,
		Path:     u.Path,
		Query:    QueryParamsFromRaw(u.RawQuery),
		Fragment: u.Fragment,
		Auth:     auths,
	}
}

func PrototypeFromString(method Method, u string) *Prototype {
	var auths auth.Auths

	rest, frag := npncore.SplitString(u, '#', true)
	if len(frag) > 0 {
		frag, _ = url.QueryUnescape(frag)
	}
	rest, query := npncore.SplitString(rest, '?', true)
	proto, rest := npncore.SplitString(rest, ':', true)
	rest = strings.TrimPrefix(strings.TrimPrefix(rest, "/"), "/")
	rest, path := npncore.SplitString(rest, '/', true)
	if len(path) > 0 {
		path, _ = url.PathUnescape(path)
	}
	aut, host := npncore.SplitString(rest, '@', true)
	if host == "" {
		host = aut
		aut = ""
	}

	if aut != "" {
		user, pass := npncore.SplitString(aut, ':', true)
		a := &auth.Basic{Username: user, Password: pass}
		auths = auth.Auths{a}
	}
	return &Prototype{
		Method:   method,
		Protocol: proto,
		Domain:   host,
		Path:     path,
		Query:    QueryParamsFromRaw(query),
		Fragment: frag,
		Auth:     auths,
	}
}
