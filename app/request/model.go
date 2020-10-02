package request

import (
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
)

type Request struct {
	Key         string     `json:"key,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Prototype   *Prototype `json:"prototype"`
}

func NewRequest() *Request {
	return &Request{Prototype: NewPrototype()}
}

func FromString(key string, content string) (*Request, error) {
	ret := &Request{}
	content = strings.TrimSpace(content)
	if strings.HasPrefix(content, `"`) || strings.HasPrefix(content, "http") {
		u := strings.TrimPrefix(strings.TrimSuffix(content, `"`), `"`)
		proto := PrototypeFromString(u)
		ret.Prototype = proto
	} else {
		errRequest := npncore.FromJSONStrict([]byte(content), ret)
		if errRequest != nil {
			proto := &Prototype{}
			errProto := npncore.FromJSONStrict([]byte(content), proto)
			if errProto != nil {
				return nil, errors.Wrap(errRequest, "unable to parse request from ["+content+"]")
			}
			ret.Prototype = proto
		}
	}
	return ret.Normalize(key), nil
}

func (r *Request) TitleWithFallback() string {
	if len(r.Title) == 0 {
		return r.Key
	}
	return r.Title
}

func (r *Request) Options() *Options {
	if r.Prototype == nil {
		return &Options{}
	}
	if r.Prototype.Options == nil {
		return &Options{}
	}
	return r.Prototype.Options
}

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

type Requests []*Request
