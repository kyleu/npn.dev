package routes

import (
	"github.com/kyleu/npn/npncore"
	"strings"
)

func Name(params ...string) string {
	return strings.Join(params, ".")
}

func Path(params ...string) string {
	ret := ""
	for _, p := range params {
		ret = ret + "/" + p
	}
	return ret
}

func Adm(params ...string) string {
	params = append([]string{npncore.KeyAdmin}, params...)
	return Path(params...)
}
