package request

import "github.com/kyleu/npn/app/auth"

type Request struct {
	Key         string     `json:"-"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Prototype   *Prototype `json:"prototype"`
}

func (r *Request) TitleWithFallback() string {
	if len(r.Title) == 0 {
		return r.Key
	}
	return r.Title
}

type Requests []*Request

var MockRequest = &Request{
	Key:         "mock",
	Title:       "Mock Request",
	Description: "a mock request",
	Prototype: &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolHTTPS,
		Domain:   "google.com",
		Path:     "search",
		Query:    QueryParams{{Key: "q", Value: "foo"}, {Key: "x", Value: "1"}},
		Fragment: "hash",
		Headers:  Headers{},
		Auth:     auth.Auths{},
		Body:     nil,
		Options: &Options{
			Timeout:               0,
			IgnoreRedirects:       false,
			IgnoreReferrer:        false,
			IgnoreCerts:           false,
			ExcludeDefaultHeaders: nil,
			ReadCookieJars:        nil,
			WriteCookieJar:        "",
			SSLCert:               "",
			UserAgentOverride:     "",
		},
	},
}

func (r *Request) Normalize(key string) *Request {
	if r == nil {
		r = &Request{}
	}
	r.Key = key
	if r.Prototype == nil {
		r.Prototype = NewPrototype()
	}
	r.Prototype = r.Prototype.Normalize()
	return r
}
