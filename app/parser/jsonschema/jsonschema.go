package parsejsonschema

import (
	"path"
	"strings"

	"emperror.dev/errors"

	parseutil "github.com/kyleu/npn/app/parser/util"
	"github.com/kyleu/npn/app/schema"
	"logur.dev/logur"
)

type JSONSchemaParser struct {
	Key    string
	logger logur.Logger
}

func NewParser(logger logur.Logger) *JSONSchemaParser {
	logger = logur.WithFields(logger, map[string]interface{}{"service": schema.OriginJSONSchema.Key})
	return &JSONSchemaParser{Key: schema.OriginJSONSchema.Key, logger: logger}
}

func (p *JSONSchemaParser) Type() schema.Origin {
	return schema.OriginJSONSchema
}

func (p *JSONSchemaParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := parseutil.GetMatchingFiles(path.Join(root, "data", "jsonschema"), "*.json")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Type: schema.OriginJSONSchema})
	}
	return ret, nil
}

func (p *JSONSchemaParser) IsValid(firstChars string) error {
	if !strings.Contains(firstChars, "{") {
		return errors.New("not JSON")
	}
	return nil
}
