package request

import (
	"fmt"
	"net/http"

	"github.com/kyleu/npn/app/session"

	"github.com/kyleu/npn/npncontroller"

	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
)

type Prototype struct {
	Method   Method                    `json:"method"`
	Protocol Protocol                  `json:"protocol"`
	Domain   string                    `json:"domain"`
	Port     int                       `json:"port,omitempty"`
	Path     string                    `json:"path,omitempty"`
	Query    npncontroller.QueryParams `json:"query,omitempty"`
	Fragment string                    `json:"fragment,omitempty"`
	Headers  header.Headers            `json:"headers,omitempty"`
	Auth     *auth.Auth                `json:"auth,omitempty"`
	Body     *body.Body                `json:"body,omitempty"`
	Options  *Options                  `json:"options,omitempty"`
}

func NewPrototype() *Prototype {
	ret := &Prototype{}
	return ret.Normalize()
}

func (p *Prototype) Normalize() *Prototype {
	if len(p.Method.Key) == 0 {
		p.Method = MethodGet
	}
	if len(p.Protocol.Key) == 0 {
		p.Protocol = ProtocolHTTPS
	}
	if p.Options == nil {
		p.Options = &Options{}
	}
	return p
}

func (p *Prototype) Host() string {
	if p.Port == 0 {
		return p.Domain
	}
	return fmt.Sprintf("%v:%v", p.Domain, p.Port)
}

func (p *Prototype) ToHTTP(sess *session.Session) *http.Request {
	ret := &http.Request{
		Method: p.Method.Key,
		URL:    p.URL(),
		Header: p.FinalHeaders(sess).ToHTTP(),
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
