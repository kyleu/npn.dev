package transform

import (
	"path"

	"emperror.dev/errors"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/npn/app/collection"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/session"
)

func OpenAPI2Import(data []byte) (*openapi3.Swagger, error) {
	x := &openapi2.Swagger{}
	err := yaml.Unmarshal(data, x)
	if err != nil {
		return nil, err
	}
	swag, err := openapi2conv.ToV3Swagger(x)
	return swag, err
}

func OpenAPI3Import(data []byte) (*openapi3.Swagger, error) {
	swag, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(data)
	return swag, err
}

func OpenAPIToFullCollection(swag *openapi3.Swagger) (*collection.FullCollection, error) {
	coll := openAPIToCollection(swag.Info)
	reqs, sess, err := openAPIToRequests(swag)
	if err != nil {
		return nil, err
	}

	ret := &collection.FullCollection{Coll: coll, Requests: reqs, Sess: sess}
	return ret, nil
}

func openAPIToRequests(swag *openapi3.Swagger) (request.Requests, *session.Session, error) {
	reqs := make(request.Requests, 0)
	sess := &session.Session{}

	url := ""
	if len(swag.Servers) > 0 {
		url = swag.Servers[0].URL
	}

	proto := request.PrototypeFromString(url)

	for k, p := range swag.Paths {
		newReqs, err := openAPIPathToRequests(k, p, proto)
		if err != nil {
			return nil, nil, err
		}
		reqs = append(reqs, newReqs...)
	}

	return reqs, sess, nil
}

func openAPIPathToRequests(pathKey string, pathItem *openapi3.PathItem, proto *request.Prototype) (request.Requests, error) {
	ret := make(request.Requests, 0)
	ops := pathItem.Operations()
	for meth, op := range ops {
		p := proto.Clone()
		p.Method = request.MethodFromString(meth)
		p.Path = path.Join(proto.Path, pathKey)

		rk := op.OperationID
		if rk == "" {
			rk = npncore.Slugify(op.Description)
		}
		if rk == "" {
			rk = npncore.Slugify(p.Method.Key + "_" + p.Path)
		}
		if rk == "" {
			return nil, errors.New("unable to determine action name")
		}

		ret = append(ret, &request.Request{
			Key:         rk,
			Title:       "",
			Description: "",
			Prototype:   p,
		})
	}
	return ret, nil
}

func openAPIToCollection(i *openapi3.Info) *collection.Collection {
	return &collection.Collection{Key: npncore.Slugify(i.Title), Title: i.Title, Description: i.Description}
}
