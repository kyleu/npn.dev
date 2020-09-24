package request

import (
	"fmt"
	"net/http"

	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
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
	}
	return fmt.Sprintf("%v:%v", p.Domain, p.Port)
}

func (p *Prototype) ToHTTP() *http.Request {
	ret := &http.Request{
		Method: p.Method.Key,
		URL:    p.URL(),
		Header: p.FinalHeaders().ToHTTP(),
		Body:   p.Body.ToHTTP(),
		Close:  false,
		Host:   p.Host(),
	}

	cl := p.ContentLength()
	if cl > 0 {
		ret.ContentLength = cl
	}

	return ret
}
