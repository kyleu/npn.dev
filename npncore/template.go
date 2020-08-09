package npncore

import (
	"bytes"
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
