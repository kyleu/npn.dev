package request

import (
	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/header"
)

func (r *Request) Normalize(key string) *Request {
	if r == nil {
		return nil
	}
	if key != "" {
		r.Key = key
	}
	if r.Key == "" {
		r.Key = "untitled-" + npncore.RandomString(6)
	}
	if r.Prototype == nil {
		r.Prototype = NewPrototype()
	}
	r.Prototype = r.Prototype.Normalize()
	return r
}

func (r *Request) Minify() *Request {
	if r.Prototype == nil {
		r.Prototype = NewPrototype()
	}
	r.Prototype = r.Prototype.Minify()
	return r
}

func (p *Prototype) Normalize() *Prototype {
	if p.Method.Key == "" {
		p.Method = MethodGet
	}
	if p.Protocol.Key == "" {
		p.Protocol = ProtocolHTTPS
	}

	p.Trim()

	if p.Options == nil {
		p.Options = &Options{}
	}
	return p
}

func (p *Prototype) Minify() *Prototype {
	p.Trim()

	if len(p.Headers) == 0 {
		p.Headers = nil
	}
	if len(p.Query) == 0 {
		p.Query = nil
	}

	if p.Body != nil && p.Body.Type == "" {
		p.Body = nil
	}
	if p.Options != nil && p.Options.Empty() {
		p.Options = nil
	}
	return p
}

func (p *Prototype) Trim() {
	qp := make(npncontroller.QueryParams, 0, len(p.Query))
	for _, x := range p.Query {
		if len(x.Key) > 0 {
			qp = append(qp, x)
		}
	}
	p.Query = qp

	headers := make(header.Headers, 0, len(p.Headers))
	for _, x := range p.Headers {
		if len(x.Key) > 0 {
			headers = append(headers, x)
		}
	}

	p.Headers = headers
}
