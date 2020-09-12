package request

import (
	"fmt"
	"net/url"
	"strings"
)

func (p *Prototype) URL() *url.URL {
	user, pass := p.Auth.GetBasic()
	var ui *url.Userinfo
	if len(user) > 0 {
		ui = url.UserPassword(user, pass)
	}
	return &url.URL{
		Scheme:   p.Protocol.String(),
		User:     ui,
		Host:     p.Domain,
		RawPath:  p.Path,
		RawQuery: p.Query.ToURL(),
		Fragment: p.Fragment,
	}
}

func (p *Prototype) URLString() string {
	domain := p.Domain
	if p.Auth.HasBasic() {
		user, pass := p.Auth.GetBasic()
		domain = fmt.Sprintf("%v:%v@%v", url.PathEscape(user), url.PathEscape(pass), p.Domain)
	}
	ret := fmt.Sprintf("%v://%v", p.Protocol.Key, domain)
	if len(p.Path) > 0 {
		ret += "/" + strings.TrimPrefix(p.Path, "/")
	}
	if len(p.Query) > 0 {
		ret += "?" + p.Query.ToURL()
	}
	if len(p.Fragment) > 0 {
		ret += "#" + url.QueryEscape(p.Fragment)
	}
	return ret
}

type URLPart struct {
	Key   string `json:"k,omitempty"`
	Value string `json:"v,omitempty"`
}

func (p *Prototype) URLParts() []*URLPart {
	ret := []*URLPart{}
	var add = func(k string, v string) {
		ret = append(ret, &URLPart{Key: k, Value: v})
	}
	add("protocol", p.Protocol.String())
	add("", "://")
	if p.Auth.HasBasic() {
		user, pass := p.Auth.GetBasic()
		add("auth", fmt.Sprintf("%v:%v", url.PathEscape(user), url.PathEscape(pass)))
		add("", "@")
	}
	add("domain", p.Domain)
	if len(p.Path) > 0 {
		add("", "/")
		add("path", strings.TrimPrefix(p.Path, "/"))
	}
	if len(p.Query) > 0 {
		add("", "?")
		add("query", p.Query.ToURL())
	}
	if len(p.Fragment) > 0 {
		add("", "#")
		add("fragment", p.Fragment)
	}
	return ret
}

func URLColor(key string) string {
	switch key {
	case "protocol":
		return "green-fg"
	case "auth":
		return "green-fg"
	case "domain":
		return "blue-fg"
	case "path":
		return "bluegrey-fg"
	case "query":
		return "purple-fg"
	default:
		return ""
	}
}
