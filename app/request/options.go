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
