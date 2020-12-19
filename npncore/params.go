package npncore

import (
	"fmt"
	"strings"

	"logur.dev/logur"
)

// A map with arbitrary string keys associated to a string array containing all allowed columns
var AllowedColumns = map[string][]string{}

// Details of a specific set of ordering parameters, with limit and offset
type Params struct {
	Key       string    `json:"key"`
	Orderings Orderings `json:"orderings,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Offset    int       `json:"offset,omitempty"`
}

// Updates or creates Params with the provided Orderings
func ParamsWithDefaultOrdering(key string, params *Params, orderings ...*Ordering) *Params {
	if params == nil {
		params = &Params{Key: key}
	}

	if len(params.Orderings) == 0 {
		params.Orderings = orderings
	}

	return params
}

// Clones this Params, replacing the orderings with the provided arguments
func (p *Params) CloneOrdering(orderings ...*Ordering) *Params {
	return &Params{Key: p.Key, Orderings: orderings, Limit: p.Limit, Offset: p.Offset}
}

// Indicates if there is more data past the provided page
func (p *Params) HasNextPage(count int) bool {
	return count > (p.Offset + p.Limit)
}

// Returns a clone of this Params, configured for the next page
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

// Indicates if there is data prior to the provided page
func (p *Params) HasPreviousPage() bool {
	return p.Offset > 0
}

// Returns a clone of this Params, configured for the previous page
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

// Returns the Orderings of this Params that match the provided column
func (p *Params) GetOrdering(col string) *Ordering {
	var ret *Ordering

	for _, o := range p.Orderings {
		if o.Column == col {
			ret = o
		}
	}

	return ret
}

// converts this Params into a SQL order by clause
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

// Filters this Params, limiting columns to those matching the AllowedColumns
func (p *Params) Filtered(logger logur.Logger) *Params {
	if len(p.Orderings) > 0 {
		allowed := make(Orderings, 0)

		for _, o := range p.Orderings {
			containsCol := false
			available, ok := AllowedColumns[p.Key]

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
				const msg = "no column [%v] for [%v] available in allowed columns [%v]"
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
		ol += fmt.Sprint(p.Limit)
	}
	ord := make([]string, 0, len(p.Orderings))
	for _, o := range p.Orderings {
		ord = append(ord, o.String())
	}
	return fmt.Sprintf("%v(%v): %v", p.Key, ol, strings.Join(ord, " / "))
}
