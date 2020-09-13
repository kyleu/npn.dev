package request

import (
	"fmt"
	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
	"net/http"
	"strconv"
	"strings"
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
		curr := p.ContentType()
		if len(curr) > 0 {
			ret = append(header.Headers{&header.Header{Key: "Content-Type", Value: curr}}, ret...)
		}
	}
	return ret
}

func (p *Prototype) ExcludesHeader(k string) bool {
	if p.Options == nil {
		return false
	}
	for _, ex := range p.Options.ExcludeDefaultHeaders {
		if strings.EqualFold(ex, k) {
			return true
		}
	}
	return false
}

func (p *Prototype) ContentType() string {
	curr := p.Headers.GetValue("Content-Type")
	if len(curr) > 0 {
		return curr
	}
	if p.Body == nil {
		return ""
	}
	return p.Body.Config.MimeType()
}

func (p *Prototype) ToHTTP() *http.Request {
	fh := p.FinalHeaders()
	cls := fh.GetValue("Content-Length")

	cl := int64(0)
	if len(cls) == 0 {
		cl = p.Body.ContentLength()
	} else {
		x, _ := strconv.Atoi(cls)
		cl = int64(x)
	}
	return &http.Request{
		Method:           p.Method.Key,
		URL:              p.URL(),
		Header:           fh.ToHTTP(),
		Body:             p.Body.ToHTTP(),
		ContentLength:    cl,
		Close:            false,
		Host:             p.Host(),
	}
}
