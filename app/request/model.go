package request

import (
	"github.com/sirupsen/logrus"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npncore"
)

type Request struct {
	Key         string     `json:"key,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Prototype   *Prototype `json:"prototype"`
}

func FromString(key string, content string) (*Request, error) {
	ret := &Request{}
	content = strings.TrimSpace(content)
	if strings.HasPrefix(content, "{") {
		b := []byte(content)
		errRequest := npncore.FromJSONStrict(b, ret)
		if errRequest != nil {
			proto := &Prototype{}
			errProto := npncore.FromJSONStrict(b, proto)
			if errProto != nil {
				return nil, errors.Wrap(errRequest, "unable to parse request from ["+content+"]")
			}
			ret.Prototype = proto
		}
	} else {
		u := strings.TrimPrefix(strings.TrimSuffix(content, `"`), `"`)
		proto := PrototypeFromString(u)
		ret.Prototype = proto
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

func (r *Request) Merge(data npncore.Data, logger *logrus.Logger) *Request {
	return &Request{
		Key:         r.Key,
		Title:       npncore.MergeLog("title", r.Title, data, logger),
		Description: npncore.MergeLog("description", r.Description, data, logger),
		Prototype:   r.Prototype.Merge(data, logger),
	}
}

type Requests []*Request
