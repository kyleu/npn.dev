package routes

import "strings"

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
