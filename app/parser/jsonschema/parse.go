package parsejsonschema

import (
	"os"

	parseutil "github.com/kyleu/npn/app/parser/util"
	"github.com/kyleu/npn/app/schema"

	"emperror.dev/errors"
	"github.com/santhosh-tekuri/jsonschema/v2"
	_ "github.com/santhosh-tekuri/jsonschema/v2/httploader" // For resolving schema refs via http
)

func (p *JSONSchemaParser) Parse(paths []string) (*parseutil.ParseResponse, error) {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginJSONSchema, Source: paths[0]}
	return p.parse(paths, parseutil.NewParseResponse(paths, md))
}

func (p *JSONSchemaParser) parse(paths []string, ret *parseutil.ParseResponse) (*parseutil.ParseResponse, error) {
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

func (p *JSONSchemaParser) parsePath(fn string, ret *parseutil.ParseResponse) (*parseutil.ParseResponse, error) {
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

	if len(js.Description) > 0 {
		ret.Schema.Description = js.Description
	}

	_, err = parseModel(ret.Schema, []string{}, "root", js)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
