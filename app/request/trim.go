package request

import (
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncore"
)

func (r *Request) Normalize(key string) *Request {
	if r == nil {
		return nil
	}
	if len(key) > 0 {
		r.Key = key
	}
	if len(r.Key) == 0 {
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
	if len(p.Method.Key) == 0 {
		p.Method = MethodGet
	}
	if len(p.Protocol.Key) == 0 {
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

	if p.Body != nil && len(p.Body.Type) == 0 {
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
