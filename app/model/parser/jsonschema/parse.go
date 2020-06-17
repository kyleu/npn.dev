package parsejsonschema

import (
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"github.com/santhosh-tekuri/jsonschema/v2"
	_ "github.com/santhosh-tekuri/jsonschema/v2/httploader"
	"emperror.dev/errors"
	"os"
)

func (p *JSONSchemaParser) ParseJSONSchemaFile(paths []string) (*JSONSchemaResponse, error) {
	return p.parse(paths, NewJSONSchemaResponse(paths))
}

func (p *JSONSchemaParser) parse(paths []string, ret *JSONSchemaResponse) (*JSONSchemaResponse, error) {
	rsp := ret
	var err error
	for _, pth := range paths {
		rsp, err = p.parsePath(pth, rsp)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing JSON schema")
		}
	}
	return rsp, nil
}

func (p *JSONSchemaParser) parsePath(fn string, ret *JSONSchemaResponse) (*JSONSchemaResponse, error) {
	reader, err := os.Open(fn)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file ["+fn+"]")
	}
	defer func() { _ = reader.Close() }()

	url := "https://npn.dev/schema"
	compiler := jsonschema.NewCompiler()
	err = compiler.AddResource(url, reader)
	if err != nil {
		return nil, errors.Wrap(err, "unable to add resource for JSON Schema file ["+fn+"]")
	}
	js, err := compiler.Compile(url)
	if err != nil {
		return nil, err
	}

	if len(js.Title) > 0 {
		ret.Schema.Title = js.Title
	}

	_, err = parseModel(ret, "root", js)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func parseModel(ret *JSONSchemaResponse, key string, js *jsonschema.Schema) (*schema.Model, error) {
	model := &schema.Model{Key: key, Type: schema.ModelTypeStruct, Metadata: nil}

	for fieldName, prop := range js.Properties {
		err := model.AddField(parseField(fieldName, prop))
		if err != nil {
			return nil, err
		}
	}

	err := ret.Schema.AddModel(model)
	if err != nil {
		return nil, err
	}
	debugSchema(ret, js)
	return model, nil
}

func parseField(key string, js *jsonschema.Schema) *schema.Field {
	return &schema.Field{Key: key, Type: schematypes.Wrap(schematypes.String{})}
}

func debugSchema(ret *JSONSchemaResponse, js *jsonschema.Schema) {
	ret.Data = append(ret.Data, map[string]interface{}{
		"format": js.Format,
		"desc":   js.Description,
		"types":  js.Types,
		"url":    js.URL,
		"props":  len(js.Properties),
	})
}

