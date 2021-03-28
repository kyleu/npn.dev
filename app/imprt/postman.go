package imprt

import (
	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
)

func parsePostman(content []byte, src string) *phase {
	c := string(content)
	pc, err := transform.PostmanImport(content)
	if err != nil {
		return errorPhase(errors.Wrap(err, "unhandled Postman "+src+" error"), c)
	}
	ret, err := transform.PostmanToFullCollection(pc)
	if err != nil {
		return errorPhase(errors.Wrap(err, "error transforming Postman "+src), content)
	}
	return &phase{Key: "postman", Value: ret, Final: true}
}
