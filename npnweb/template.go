package npnweb

import "fmt"

var IconContent = func(color string) string {
	return ""
}

var NavbarContent = func(color string) string {
	return ""
}

func AdminLink(params ...string) string {
	ret := fmt.Sprintf("admin")
	for _, p := range params {
		ret = ret + "." + p
	}

	return ret
}
