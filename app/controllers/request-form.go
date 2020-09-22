package controllers

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/app/body"
	"github.com/kyleu/npn/app/header"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/npncore"
)

type prototypeForm struct {
	Method  string `mapstructure:"method"`
	URL     string `mapstructure:"url"`
	Headers string `mapstructure:"headers"`
	Auth    string `mapstructure:"auth"`
	Body    string `mapstructure:"body"`
	Options string `mapstructure:"options"`
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

	o := &request.Options{}
	err = npncore.FromJSON([]byte(f.Options), o)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse options")
	}
	proto.Options = o

	req := &request.Request{
		Key:         f.Key,
		Title:       f.Title,
		Description: f.Description,
		Prototype:   proto,
	}

	return req.Normalize(f.Key), nil
}
