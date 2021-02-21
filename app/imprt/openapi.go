package imprt

import (
	"github.com/kyleu/npn/app/transform"

	"emperror.dev/errors"
)

func parseOpenAPI3(content []byte, src string) *phase {
	c := string(content)
	oapi, err := transform.OpenAPI3Import(content)
	if err != nil {
		return errorPhase(errors.Wrap(err, "unhandled OpenAPI 3 " + src + " error"), c)
	}
	ret, err := transform.OpenAPIToFullCollection(oapi)
	if err != nil {
		return errorPhase(errors.Wrap(err, "error transforming OpenAPI 3 " + src), content)
	}
	return &phase{Key: "openapi", Value: ret, Final: true}
}

func parseOpenAPI2(content []byte, src string) *phase {
	c := string(content)
	oapi, err := transform.OpenAPI2Import(content)
	if err != nil {
		return errorPhase(errors.Wrap(err, "unhandled OpenAPI 2 " + src + " error"), c)
	}
	ret, err := transform.OpenAPIToFullCollection(oapi)
	if err != nil {
		return errorPhase(errors.Wrap(err, "error transforming OpenAPI 2 " + src), content)
	}
	return &phase{Key: "openapi", Value: ret, Final: true}
}
