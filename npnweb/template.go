package npnweb

import "fmt"

var IconContent = ""

func AdminLink(params ...string) string {
	ret := fmt.Sprintf("admin")
	for _, p := range params {
		ret = ret + "." + p
	}

	return ret
}
