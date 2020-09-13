package request

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/kyleu/npn/app/auth"
	"github.com/kyleu/npn/npncore"
)

func PrototypeFromURL(u *url.URL) *Prototype {
	var auths auth.Auths
	if u.User != nil {
		p, _ := u.User.Password()
		a := auth.NewBasic(u.User.Username(), p, false)
		auths = auth.Auths{a}
	}
	domain, portString := npncore.SplitString(u.Host, ':', true)

	port := 0
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(u.Scheme),
		Domain:   domain,
		Port:     port,
		Path:     u.Path,
		Query:    QueryParamsFromRaw(u.RawQuery),
		Fragment: u.Fragment,
		Auth:     auths,
	}
}

func PrototypeFromString(u string) *Prototype {
	var auths auth.Auths

	rest, frag := npncore.SplitString(u, '#', true)
	if len(frag) > 0 {
		frag, _ = url.QueryUnescape(frag)
	}
	rest, query := npncore.SplitString(rest, '?', true)
	proto, rest := npncore.SplitString(rest, ':', true)
	rest = strings.TrimPrefix(strings.TrimPrefix(rest, "/"), "/")
	rest, path := npncore.SplitString(rest, '/', true)
	if len(path) > 0 {
		path, _ = url.PathUnescape(path)
	}
	aut, host := npncore.SplitString(rest, '@', true)
	if host == "" {
		host = aut
		aut = ""
	}
	host, portString := npncore.SplitString(host, ':', true)
	port := 0
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	if aut != "" {
		user, pass := npncore.SplitString(aut, ':', true)
		a := auth.NewBasic(user, pass, false)
		auths = auth.Auths{a}
	}
	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(proto),
		Domain:   host,
		Port:     port,
		Path:     path,
		Query:    QueryParamsFromRaw(query),
		Fragment: frag,
		Auth:     auths,
	}
}
