package request

import "github.com/kyleu/npn/app/auth"

type Request struct {
	Key         string     `json:"key"`
	Description string     `json:"description,omitempty"`
	Prototype   *Prototype `json:"prototype"`
}

var MockRequest = &Request{
	Key:         "mock",
	Description: "a mock request",
	Prototype: &Prototype{
		Method:   MethodGet,
		Protocol: "https",
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
