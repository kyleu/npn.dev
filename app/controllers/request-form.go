package controllers

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
	"strconv"
)

type optionsForm struct {
	Timeout               string `mapstructure:"opt.timeout"`
	IgnoreRedirects       string `mapstructure:"opt.ignoreRedirects"`
	IgnoreReferrer        string `mapstructure:"opt.ignoreReferrer"`
	IgnoreCerts           string `mapstructure:"opt.ignoreCerts"`
	ExcludeDefaultHeaders string `mapstructure:"opt.excludeDefaultHeaders"`
	ReadCookieJars        string `mapstructure:"opt.readCookieJars"`
	WriteCookieJar        string `mapstructure:"opt.writeCookieJar"`
	SSLCert               string `mapstructure:"opt.sslCert"`
	UserAgentOverride     string `mapstructure:"opt.userAgentOverride"`
}

type prototypeForm struct {
	Method      string `mapstructure:"method"`
	URL         string `mapstructure:"url"`
	Headers     string `mapstructure:"headers"`
	Auth        string `mapstructure:"auth"`
	Body        string `mapstructure:"body"`
	optionsForm `mapstructure:",squash"`
}

type requestForm struct {
	Coll          string `mapstructure:"coll"`
	OriginalKey   string `mapstructure:"originalKey"`
	Key           string `mapstructure:"key"`
	Title         string `mapstructure:"title"`
	Description   string `mapstructure:"description"`
	prototypeForm `mapstructure:",squash"`
}

func (f *requestForm) ToRequest() (*request.Request, error) {
	proto := request.PrototypeFromString(f.URL)

	h := &header.Headers{}
	err := npncore.FromJSON([]byte(f.Headers), h)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse headers")
	}
	proto.Headers = *h

	a := &auth.Auths{}
	err = npncore.FromJSON([]byte(f.Auth), a)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse auth")
	}
	proto.Auth = *a

	b := &body.Body{}
	err = npncore.FromJSON([]byte(f.Body), b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse body")
	}
	if len(b.Type) > 0 {
		proto.Body = b
	}

	proto.Options, err = parseOptions(f)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse options")
	}

	req := &request.Request{
		Key:         f.Key,
		Title:       f.Title,
		Description: f.Description,
		Prototype:   proto,
	}

	return req.Normalize(f.Key), nil
}

func parseOptions(f *requestForm) (*request.Options, error) {
	timeout, err := strconv.Atoi(f.Timeout)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse numeric timeout from [" + f.Timeout + "]")
	}

	excludeDefaultHeaders := &[]string{}
	err = npncore.FromJSON([]byte(f.ExcludeDefaultHeaders), excludeDefaultHeaders)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse exclusions from [" + f.ExcludeDefaultHeaders + "]")
	}

	readCookieJars := &[]string{}
	err = npncore.FromJSON([]byte(f.ReadCookieJars), readCookieJars)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse cookie jars from [" + f.ReadCookieJars + "]")
	}

	return &request.Options{
		Timeout:               timeout,
		IgnoreRedirects:       f.IgnoreRedirects == "true",
		IgnoreReferrer:        f.IgnoreReferrer == "true",
		IgnoreCerts:           f.IgnoreCerts == "true",
		ExcludeDefaultHeaders: *excludeDefaultHeaders,
		ReadCookieJars:        *readCookieJars,
		WriteCookieJar:        f.WriteCookieJar,
		SSLCert:               f.SSLCert,
		UserAgentOverride:     f.UserAgentOverride,
	}, nil
}
