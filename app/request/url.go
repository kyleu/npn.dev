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
		Host:     p.Host(),
		RawPath:  p.Path,
		RawQuery: p.Query.ToURL(),
		Fragment: p.Fragment,
	}
}

func (p *Prototype) FullPathString() string {
	trimmed := strings.TrimSpace(strings.TrimPrefix(p.Path, "/"))
	if len(trimmed) == 0 {
		return ""
	}
	ret := "/" + trimmed
	if len(p.Query) > 0 {
		ret += "?" + p.Query.ToURL()
	}
	if len(p.Fragment) > 0 {
		ret += "#" + url.QueryEscape(p.Fragment)
	}
	return ret
}

func (p *Prototype) URLString() string {
	domain := p.Host()
	if p.Auth.HasBasic() {
		user, pass := p.Auth.GetBasic()
		domain = fmt.Sprintf("%v:%v@%v", url.PathEscape(user), url.PathEscape(pass), p.Host())
	}
	ret := fmt.Sprintf("%v://%v", p.Protocol.Key, domain)
	ret += p.FullPathString()
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
		add("username", url.PathEscape(user))
		add("", ":")
		add("password", url.PathEscape(pass))
		add("", "@")
	}
	add("domain", p.Domain)
	if p.Port > 0 {
		add("", ":")
		add("port", fmt.Sprintf("%v", p.Port))
	}
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
	case "protocol", "auth", "username", "password":
		return "green-fg"
	case "domain", "port":
		return "blue-fg"
	case "path":
		return "bluegrey-fg"
	case "query":
		return "purple-fg"
	default:
		return ""
	}
}
