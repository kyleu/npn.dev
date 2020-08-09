package npncore

import (
	"fmt"
	"strings"

	"logur.dev/logur"
)

var allowedColumns = map[string][]string{}

type Params struct {
	Key       string
	Orderings Orderings
	Limit     int
	Offset    int
}

func ParamsWithDefaultOrdering(key string, params *Params, orderings ...*Ordering) *Params {
	if params == nil {
		params = &Params{Key: key}
	}

	if len(params.Orderings) == 0 {
		params.Orderings = orderings
	}

	return params
}

func (p *Params) CloneOrdering(orderings ...*Ordering) *Params {
	return &Params{Key: p.Key, Orderings: orderings, Limit: p.Limit, Offset: p.Offset}
}

func (p *Params) HasNextPage(count int) bool {
	return count > (p.Offset + p.Limit)
}

func (p *Params) NextPage() *Params {
	limit := p.Limit
	if limit == 0 {
		limit = 100
	}
	offset := p.Offset + limit
	if offset < 0 {
		offset = 0
	}
	return &Params{Key: p.Key, Orderings: p.Orderings, Limit: p.Limit, Offset: offset}
}

func (p *Params) HasPreviousPage() bool {
	return p.Offset > 0
}

func (p *Params) PreviousPage() *Params {
	limit := p.Limit
	if limit == 0 {
		limit = 100
	}
	offset := p.Offset - limit
	if offset < 0 {
		offset = 0
	}
	return &Params{Key: p.Key, Orderings: p.Orderings, Limit: p.Limit, Offset: offset}
}

func (p *Params) GetOrdering(col string) *Ordering {
	var ret *Ordering

	for _, o := range p.Orderings {
		if o.Column == col {
			ret = o
		}
	}

	return ret
}

func (p *Params) OrderByString() string {
	var ret = make([]string, 0, len(p.Orderings))

	for _, o := range p.Orderings {
		dir := ""
		if !o.Asc {
			dir = " desc"
		}
		snake := ToSnakeCase(o.Column)
		ret = append(ret, snake+dir)
	}

	return strings.Join(ret, ", ")
}

func (p *Params) Filtered(logger logur.Logger) *Params {
	if len(p.Orderings) > 0 {
		allowed := make(Orderings, 0)

		for _, o := range p.Orderings {
			containsCol := false
			available, ok := allowedColumns[p.Key]

			if !ok {
				logger.Warn("no columns available for [" + p.Key + "]")
			}

			for _, c := range available {
				if c == o.Column {
					containsCol = true
				}
			}

			if containsCol {
				allowed = append(allowed, o)
			} else {
				msg := "no column [%v] for [%v] available in allowed columns [%v]"
				logger.Warn(fmt.Sprintf(msg, o.Column, p.Key, OxfordComma(available, "and")))
			}
		}

		return &Params{Key: p.Key, Orderings: allowed, Limit: p.Limit, Offset: p.Offset}
	}

	return p
}

func (p *Params) String() string {
	ol := ""
	if p.Offset > 0 {
		ol += fmt.Sprintf("%v/", p.Offset)
	}
	if p.Limit > 0 {
		ol += fmt.Sprintf("%v", p.Limit)
	}
	ord := make([]string, 0, len(p.Orderings))
	for _, o := range p.Orderings {
		ord = append(ord, o.String())
	}
	return fmt.Sprintf("%v(%v): %v", p.Key, ol, strings.Join(ord, " / "))
}
