package request

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/session"

	"github.com/kyleu/libnpn/npncontroller"

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

func (p *Prototype) GetCookies() header.Cookies {
	cookieHeader := p.Headers.Get("Cookie")
	if cookieHeader == nil {
		return header.Cookies{}
	}
	h := http.Header{}
	h.Add(cookieHeader.Key, cookieHeader.Value)
	r := http.Request{Header: h}
	cooks := r.Cookies()
	ret := header.Cookies{}
	for _, c := range cooks {
		ret = append(ret, header.NewCookie(c))
	}
	return ret
}

func (p *Prototype) SetCookies(cookies header.Cookies) {
	cookies = cookies.Qualifying(p.URL())
	if len(cookies) > 0 {
		p.Headers = p.Headers.Set("Cookie", cookies.String())
	}
}

func (p *Prototype) Merge(data npncore.Data, logger *logrus.Logger) *Prototype {
	meth := p.Method
	if npncore.MergeNeeded(meth.Key) {
		meth = Method{Key: npncore.MergeLog("proto.method", meth.Key, data, logger)}
	}
	prot := p.Protocol
	if npncore.MergeNeeded(prot.Key) {
		prot = Protocol{Key: npncore.MergeLog("proto.protocol", prot.Key, data, logger)}
	}

	return &Prototype{
		Method:   meth,
		Protocol: prot,
		Domain:   npncore.MergeLog("proto.domain", p.Domain, data, logger),
		Port:     p.Port,
		Path:     npncore.MergeLog("proto.path", p.Path, data, logger),
		Query:    p.Query.Merge(data, logger),
		Fragment: npncore.MergeLog("proto.fragment", p.Fragment, data, logger),
		Headers:  p.Headers.Merge(data, logger),
		Auth:     p.Auth.Merge(data, logger),
		Body:     p.Body.Merge(data, logger),
		Options:  p.Options.Merge(data, logger),
	}
}

func (p *Prototype) Clone() *Prototype {
	return &Prototype{
		Method:   p.Method,
		Protocol: p.Protocol,
		Domain:   p.Domain,
		Port:     p.Port,
		Path:     p.Path,
		Query:    p.Query.Clone(),
		Fragment: p.Fragment,
		Headers:  p.Headers.Clone(),
		Auth:     p.Auth.Clone(),
		Body:     p.Body.Clone(),
		Options:  p.Options.Clone(),
	}
}
