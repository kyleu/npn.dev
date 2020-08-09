package parseintellij

import (
	"path"
	"strings"

	"emperror.dev/errors"

	parseutil "github.com/kyleu/npn/app/parser/util"
	"github.com/kyleu/npn/app/schema"
	"logur.dev/logur"
)

type IntelliJParser struct {
	Key    string
	logger logur.Logger
}

func NewParser(logger logur.Logger) *IntelliJParser {
	logger = logur.WithFields(logger, map[string]interface{}{"service": schema.OriginDatabase.Key})
	return &IntelliJParser{Key: schema.OriginDatabase.Key, logger: logger}
}

func (p *IntelliJParser) Type() schema.Origin {
	return schema.OriginDatabase
}

func (p *IntelliJParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := parseutil.GetMatchingFiles(path.Join(root, ".idea", "dataSources"), "*.xml")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Type: schema.OriginDatabase})
	}
	return ret, nil
}

func (p *IntelliJParser) IsValid(firstChars string) error {
	if !strings.Contains(firstChars, "<") {
		return errors.New("not XML")
	}
	return nil
}
