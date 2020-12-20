package routes

import (
	"strings"

	"github.com/kyleu/npn/npncore"
)

// Joins the arguments with "."
func Name(params ...string) string {
	return strings.Join(params, ".")
}

// Joins the arguments with "/"
func Path(params ...string) string {
	ret := ""
	for _, p := range params {
		ret = ret + "/" + p
	}
	return ret
}

// Joins the arguments with ".", prepends "admin."
func Adm(params ...string) string {
	params = append([]string{npncore.KeyAdmin}, params...)
	return Path(params...)
}
