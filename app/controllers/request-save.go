package controllers

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"
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
	proto.Body = b

	proto.Options = &request.Options{
		Timeout:               0,
		IgnoreRedirects:       f.IgnoreRedirects == "true",
		IgnoreReferrer:        f.IgnoreReferrer == "true",
		IgnoreCerts:           f.IgnoreCerts == "true",
		ExcludeDefaultHeaders: nil,
		ReadCookieJars:        nil,
		WriteCookieJar:        f.WriteCookieJar,
		SSLCert:               f.SSLCert,
		UserAgentOverride:     f.UserAgentOverride,
	}

	req := &request.Request{
		Key:         f.Key,
		Title:       f.Title,
		Description: f.Description,
		Prototype:   proto,
	}

	return req, nil
}

func RequestSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		frm := &requestForm{}
		err := npnweb.Decode(r, frm, ctx.Logger)
		if err != nil {
			return npncontroller.EResp(err)
		}
		req, err := frm.ToRequest()
		if err != nil {
			return npncontroller.EResp(err, "unable to parse request")
		}

		csvc := app.Svc(ctx.App).Collection

		err = csvc.SaveRequest(frm.Coll, req)
		if err != nil {
			return npncontroller.EResp(err, "unable to save ["+frm.Coll+"/"+req.Key+"]")
		}

		if frm.Key != frm.OriginalKey {
			err = csvc.DeleteRequest(frm.Coll, frm.OriginalKey)
			if err != nil {
				return npncontroller.EResp(err, "unable to delete ["+frm.Coll+"/"+frm.OriginalKey+"]")
			}
		}

		msg := "saved request [" + req.Key + "]"
		rt := ctx.Route(KeyRequest, "c", frm.Coll, npncore.KeyKey, req.Key)
		return npncontroller.FlashAndRedir(true, msg, rt, w, r, ctx)
	})
}
