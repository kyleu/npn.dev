package request

import (
	"github.com/kyleu/libnpn/npncore"
	"logur.dev/logur"
)

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

func (o *Options) Merge(data npncore.Data, logger logur.Logger) *Options {
	if o == nil {
		return nil
	}
	edh := make([]string, 0, len(o.ExcludeDefaultHeaders))
	for _, dh := range o.ExcludeDefaultHeaders {
		edh = append(edh, npncore.MergeLog("options.excludeDefaultHeaders", dh, data, logger))
	}
	rcj := make([]string, 0, len(o.ReadCookieJars))
	for _, cj := range o.ReadCookieJars {
		rcj = append(rcj, npncore.MergeLog("options.readCookieJars", cj, data, logger))
	}

	return &Options{
		Timeout:               o.Timeout,
		IgnoreRedirects:       o.IgnoreRedirects,
		IgnoreReferrer:        o.IgnoreReferrer,
		IgnoreCerts:           o.IgnoreCerts,
		IgnoreCookies:         o.IgnoreCookies,
		ExcludeDefaultHeaders: edh,
		ReadCookieJars:        rcj,
		WriteCookieJar:        npncore.MergeLog("options.writeCookieJar", o.WriteCookieJar, data, logger),
		SSLCert:               npncore.MergeLog("options.sslCert", o.SSLCert, data, logger),
		UserAgentOverride:     npncore.MergeLog("options.userAgentOverride", o.UserAgentOverride, data, logger),
	}
}

func (o *Options) Clone() *Options {
	if o == nil {
		return nil
	}
	return &Options{
		Timeout:               o.Timeout,
		IgnoreRedirects:       o.IgnoreRedirects,
		IgnoreReferrer:        o.IgnoreReferrer,
		IgnoreCerts:           o.IgnoreCerts,
		IgnoreCookies:         o.IgnoreCookies,
		ExcludeDefaultHeaders: o.ExcludeDefaultHeaders,
		ReadCookieJars:        o.ReadCookieJars,
		WriteCookieJar:        o.WriteCookieJar,
		SSLCert:               o.SSLCert,
		UserAgentOverride:     o.UserAgentOverride,
	}
}
