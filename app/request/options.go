package request

type Options struct {
	Timeout               int      `json:"timeout,omitempty"`
	IgnoreRedirects       bool     `json:"ignoreRedirects,omitempty"`
	IgnoreReferrer        bool     `json:"ignoreReferrer,omitempty"`
	IgnoreCerts           bool     `json:"ignoreCerts,omitempty"`
	IgnoreCookies         bool     `json:"ignoreCookies,omitempty"`
	ExcludeDefaultHeaders []string `json:"excludeDefaultHeaders,omitempty"`
	ReadCookieJars        []string `json:"readCookieJars,omitempty"`
	WriteCookieJar        string   `json:"writeCookieJar,omitempty"`
	SSLCert               string   `json:"sslCert,omitempty"`
	UserAgentOverride     string   `json:"userAgentOverride,omitempty"`
}

func (o *Options) Empty() bool {
	if o.Timeout != 0 {
		return false
	}
	if o.IgnoreRedirects {
		return false
	}
	if o.IgnoreReferrer {
		return false
	}
	if o.IgnoreCerts {
		return false
	}
	if o.IgnoreCookies {
		return false
	}
	if len(o.ExcludeDefaultHeaders) > 0 {
		return false
	}
	if len(o.ReadCookieJars) > 0 {
		return false
	}
	if len(o.WriteCookieJar) > 0 {
		return false
	}
	if len(o.SSLCert) > 0 {
		return false
	}
	if len(o.UserAgentOverride) > 0 {
		return false
	}
	return true
}
