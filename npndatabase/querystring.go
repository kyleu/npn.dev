package npndatabase

import (
	"fmt"
	"net/url"
)

func (p *Params) ToQueryString(u *url.URL) string {
	if p == nil {
		return ""
	}

	if u == nil {
		return ""
	}

	var ret = u.Query()

	delete(ret, p.Key+".o")
	delete(ret, p.Key+".l")
	delete(ret, p.Key+".x")

	for _, o := range p.Orderings {
		s := o.Column

		if !o.Asc {
			s += ".d"
		}

		ret.Add(p.Key+".o", s)
	}

	if p.Limit > 0 {
		ret.Add(p.Key+".l", fmt.Sprintf("%v", p.Limit))
	}

	if p.Offset > 0 {
		ret.Add(p.Key+".x", fmt.Sprintf("%v", p.Offset))
	}

	return ret.Encode()
}
