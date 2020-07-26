package query

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kyleu/npn/npncore"
	"strings"

	"logur.dev/logur"
)

type Params struct {
	Key       string
	Orderings Orderings
	Limit     int
	Offset    int
}

func (p *Params) Clone(orderings ...*Ordering) *Params {
	return &Params{Key: p.Key, Orderings: orderings, Limit: p.Limit, Offset: p.Offset}
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
		snake := strings.ToLower(strcase.ToSnake(o.Column))
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
				logger.Warn(fmt.Sprintf(msg, o.Column, p.Key, npncore.OxfordComma(available, "and")))
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
