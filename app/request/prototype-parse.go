package request

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/kyleu/libnpn/npncontroller"

	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/auth"
)

func PrototypeFromURL(u *url.URL) *Prototype {
	var at *auth.Auth
	if u.User != nil {
		p, _ := u.User.Password()
		a := auth.NewBasic(u.User.Username(), p, false)
		at = a
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
		Query:    npncontroller.QueryParamsFromRaw(u.RawQuery),
		Fragment: u.Fragment,
		Auth:     at,
	}
}

func PrototypeFromString(u string) *Prototype {
	var at *auth.Auth

	rest, frag := npncore.SplitString(u, '#', true)
	if len(frag) > 0 {
		frag, _ = url.QueryUnescape(frag)
	}
	rest, query := npncore.SplitString(rest, '?', true)
	proto, rest := npncore.SplitString(rest, ':', true)
	if rest == "" {
		rest = proto
		proto = "http"
	}
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
		at = a
	}
	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(proto),
		Domain:   host,
		Port:     port,
		Path:     path,
		Query:    npncontroller.QueryParamsFromRaw(query),
		Fragment: frag,
		Auth:     at,
	}
}
