package request

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kyleu/npn/npncore"

	"github.com/kyleu/npn/app/header"
)

func (p *Prototype) FinalHeaders() header.Headers {
	ret := p.Headers.Clone()

	var check = func(k string, f func() string) {
		if (!p.Headers.Contains(k)) && (!p.ExcludesHeader(k)) {
			s := f()
			if len(s) > 0 {
				h := &header.Header{Key: k, Value: s}
				ret = append(ret, h)
			}
		}
	}

	check("Accept", func() string { return "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8" })
	check("Accept-Encoding", func() string { return "gzip, deflate" })
	check("Connection", func() string { return "keep-alive" })
	check("Content-Type", p.ContentType)
	check("Content-Length", func() string {
		cl := p.Body.ContentLength()
		if cl > 0 {
			return fmt.Sprintf("%v", cl)
		}
		return ""
	})
	check("Host", p.Host)
	check("Origin", func() string {
		if p.Method == MethodPost {
			return p.Protocol.Key + "://" + p.Host()
		}
		return ""
	})
	check("User-Agent", func() string { return npncore.AppName })

	return ret
}

func (p *Prototype) ExcludesHeader(k string) bool {
	if p.Options == nil {
		return false
	}
	for _, ex := range p.Options.ExcludeDefaultHeaders {
		if strings.EqualFold(ex, k) {
			return true
		}
	}
	return false
}

func (p *Prototype) ContentType() string {
	curr := p.Headers.GetValue("Content-Type")
	if len(curr) > 0 {
		return curr
	}
	if p.Body == nil {
		return ""
	}
	return p.Body.Config.MimeType()
}

func (p *Prototype) ContentLength() int64 {
	curr := p.Headers.GetValue("Content-Length")
	if len(curr) > 0 {
		x, err := strconv.Atoi(curr)
		if err == nil {
			return int64(x)
		}
	}
	if p.Body == nil {
		return -1
	}
	return p.Body.ContentLength()
}
