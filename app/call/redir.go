package call

import (
	"path/filepath"
	"strings"

	"github.com/kyleu/npn/app/request"
)

func getRedir(rsp *Response, p *request.Prototype) *request.Prototype {
	loc := ""
	if rsp != nil {
		loc = rsp.Headers.GetValue("location")
	}
	if loc == "" {
		return nil
	}
	if strings.HasPrefix(loc, "//") {
		loc = p.Protocol.Key + ":" + loc
	}
	if !strings.Contains(loc, "://") {
		if !strings.HasPrefix(loc, "/") {
			loc = filepath.Dir(p.Path) + "/" + loc
		}
		loc = p.Protocol.Key + "://" + p.Host() + loc
	}
	redirP := request.PrototypeFromString(loc)
	redirP.Auth = p.Auth
	redirP.Headers = p.Headers
	redirP.Options = p.Options

	return redirP
}
