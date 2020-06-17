package parsejsonschema

import (
	"fmt"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"path"
)

type JSONSchemaParser struct {
	Key     string
	logger  logur.Logger
}

func NewParser(logger logur.Logger) *JSONSchemaParser {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: schema.OriginJSONSchema.Key})
	return &JSONSchemaParser{Key: schema.OriginJSONSchema.Key, logger: logger}
}

func (p *JSONSchemaParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := util.GetMatchingFiles(path.Join(root, "data", "jsonschema"), "*.json")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Origin: schema.OriginJSONSchema})
	}
	return ret, nil
}

func (p *JSONSchemaParser) log(rsp *JSONSchemaResponse, err error) {
	if err != nil {
		rsp.Schema.Errors = append(rsp.Schema.Errors, err.Error())
		p.logger.Error(fmt.Sprintf("unable to parse JSON schema: %+v", err))
	}
}
