package npncore

import (
	"bytes"
	"strings"
	"text/template"

	"emperror.dev/errors"
)

func Template(content string, arg interface{}) (string, error) {
	tmpl, err := template.New("adhoc").Parse(content)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse template from content ["+content+"]")
	}
	out := &bytes.Buffer{}
	err = tmpl.Execute(out, arg)
	if err != nil {
		return "", errors.Wrap(err, "unable to execute template")
	}
	return out.String(), nil
}

func SplitString(s string, sep byte, cutc bool) (string, string) {
	i := strings.IndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}

func SplitStringLast(s string, sep byte, cutc bool) (string, string) {
	i := strings.LastIndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}
