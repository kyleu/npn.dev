package util

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func OxfordComma(names []string, clause string) string {
	ret := ""
	for idx, name := range names {
		if idx > 0 {
			if len(clause) > 0 && idx == (len(names)-1) {
				if idx > 1 {
					ret += ","
				}
				ret += " " + clause + " "
			} else {
				ret += ", "
			}
		}
		ret += name
	}
	return ret
}

func SplitPackage(s string) ([]string, string) {
	sp := strings.Split(s, ".")
	pkg := sp[0:len(sp) - 1]
	n := sp[len(sp) - 1]
	return pkg, n
}
