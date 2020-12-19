package npncore

import (
	"fmt"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Converts a string to snake_case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Extracts the domain from an email address
func GetDomain(email string) string {
	var idx = strings.LastIndex(email, "@")
	if idx == -1 {
		return email
	}
	return email[idx:]
}

// Addes commas like strings.Join(_, ", "), along with a clause like "and", producing "a, b, and c"
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

// Converts provided array elements to strings, then joins them as a list
func ValueStrings(values []interface{}) string {
	ret := make([]string, 0, len(values))
	for _, v := range values {
		ret = append(ret, fmt.Sprintf(`"%v"`, v))
	}
	return strings.Join(ret, ", ")
}

var re *regexp.Regexp

// Extracts path parameters from the provided string
func PathParams(s string) []string {
	if re == nil {
		re = regexp.MustCompile("{([^}]*)}")
	}

	matches := re.FindAll([]byte(s), -1)

	ret := make([]string, 0, len(matches))
	for _, m := range matches {
		ret = append(ret, string(m))
	}

	return ret
}

// Limits a string to a max length, adding "..." if truncated
func TruncateString(x interface{}, max int) string {
	s := fmt.Sprint(x)
	if len(s) > max {
		return s[0:max] + "..."
	}
	return s
}
