package request

import (
	"fmt"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/npncore"
)

type Prototype struct {
	Method   Method         `json:"method"`
	Protocol Protocol       `json:"protocol"`
	Domain   string         `json:"domain"`
	Port     int            `json:"port,omitempty"`
	Path     string         `json:"path,omitempty"`
	Query    QueryParams    `json:"query,omitempty"`
	Fragment string         `json:"fragment,omitempty"`
	Headers  header.Headers `json:"headers,omitempty"`
	Auth     auth.Auths     `json:"auth,omitempty"`
	Body     *body.Body     `json:"body,omitempty"`
	Options  *Options       `json:"options,omitempty"`
}

func NewPrototype() *Prototype {
	ret := &Prototype{}
	ret = ret.Normalize()
	return ret
}

func (p *Prototype) Normalize() *Prototype {
	if len(p.Method.Key) == 0 {
		p.Method = MethodGet
	}
	if len(p.Protocol.Key) == 0 {
		p.Protocol = ProtocolHTTPS
	}
	return p
}

func (p *Prototype) Host() string {
	if p.Port == 0 {
		return p.Domain
	} else {
		return fmt.Sprintf("%v:%v", p.Domain, p.Port)
	}
}

func (p *Prototype) FinalHeaders() header.Headers {
	ret := make(header.Headers, len(p.Headers))
	for i, h := range p.Headers {
		ret[i] = h
	}
	if (!p.Headers.Contains("Host")) && (!p.ExcludesHeader("Host")) {
		host := &header.Header{Key: "Host", Value: p.Host()}
		ret = append(header.Headers{host}, ret...)
	}
	if (!p.Headers.Contains("Content-Type")) && (!p.ExcludesHeader("Content-Type")) {
		ct := p.ContentType()
		if len(ct) > 0 {
			ct := &header.Header{Key: "Content-Type", Value: ct}
			ret = append(header.Headers{ct}, ret...)
		}
	}
	return ret
}

func (p *Prototype) ExcludesHeader(k string) bool {
	k = strings.ToLower(k)
	if p.Options == nil {
		return false
	}
	for _, ex := range p.Options.ExcludeDefaultHeaders {
		if ex == k {
			return true
		}
	}
	return false
}


func (p *Prototype) ContentType() string {
	if p.Body == nil {
		return ""
	}
	switch p.Body.Type {
	case body.KeyTemp:
		return p.Body.Config.MimeType()
	default:
		return "text/html"
	}
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
		Host:             p.Host(),
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

func PrototypeFromURL(u *url.URL) *Prototype {
	var auths auth.Auths
	if u.User != nil {
		p, _ := u.User.Password()
		a := auth.NewBasic(u.User.Username(), p, false)
		auths = auth.Auths{a}
	}
	domain, portString := npncore.SplitString(u.Host, ':', true)

	port := 0
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(u.Scheme),
		Domain:   domain,
		Port:     port,
		Path:     u.Path,
		Query:    QueryParamsFromRaw(u.RawQuery),
		Fragment: u.Fragment,
		Auth:     auths,
	}
}

func PrototypeFromString(u string) *Prototype {
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
	host, portString := npncore.SplitString(host, ':', true)
	port := 0
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	if aut != "" {
		user, pass := npncore.SplitString(aut, ':', true)
		a := auth.NewBasic(user, pass, false)
		auths = auth.Auths{a}
	}
	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(proto),
		Domain:   host,
		Port:     port,
		Path:     path,
		Query:    QueryParamsFromRaw(query),
		Fragment: frag,
		Auth:     auths,
	}
}
