package npnweb

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func ParamSetFromRequest(r *http.Request) npncore.ParamSet {
	ret := make(npncore.ParamSet)
	for qk, qs := range r.URL.Query() {
		if strings.Contains(qk, ".") {
			for _, qv := range qs {
				ret = apply(ret, qk, qv)
			}
		}
	}
	return ret
}

func apply(ps npncore.ParamSet, qk string, qv string) npncore.ParamSet {
	switch {
	case strings.HasSuffix(qk, ".o"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".o"))
		asc := true
		if strings.HasSuffix(qv, ".d") {
			asc = false
			qv = qv[0 : len(qv)-2]
		}
		curr.Orderings = append(curr.Orderings, &npncore.Ordering{Column: qv, Asc: asc})
	case strings.HasSuffix(qk, ".l"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".l"))
		li, err := strconv.ParseInt(qv, 10, 64)
		if err == nil {
			curr.Limit = int(li)
			max := 10000
			if curr.Limit > max {
				curr.Limit = max
			}
		}
	case strings.HasSuffix(qk, ".x"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".x"))
		xi, err := strconv.ParseInt(qv, 10, 64)
		if err == nil {
			curr.Offset = int(xi)
		}
	}
	return ps
}

func getCurr(q npncore.ParamSet, key string) *npncore.Params {
	curr, ok := q[key]
	if !ok {
		curr = &npncore.Params{Key: key}
		q[key] = curr
	}
	return curr
}

func IDFromParams(key string, m map[string]string) (*uuid.UUID, error) {
	retOut, ok := m[npncore.KeyID]
	if !ok {
		return nil, errors.New("params do not contain \"id\"")
	}

	ret, err := uuid.FromString(retOut)
	return &ret, err
}
