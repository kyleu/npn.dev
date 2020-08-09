package parseprotobuf

import (
	"fmt"
	"path"
	"strings"

	"emperror.dev/errors"

	parseutil "github.com/kyleu/npn/app/parser/util"
	"github.com/kyleu/npn/app/schema"
	"logur.dev/logur"
)

type ProtobufParser struct {
	Key    string
	logger logur.Logger
}

func NewParser(logger logur.Logger) *ProtobufParser {
	logger = logur.WithFields(logger, map[string]interface{}{"service": schema.OriginProtobuf.Key})
	return &ProtobufParser{Key: schema.OriginProtobuf.Key, logger: logger}
}

func (p *ProtobufParser) Type() schema.Origin {
	return schema.OriginProtobuf
}

func (p *ProtobufParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := parseutil.GetMatchingFiles(path.Join(root, "data", "protobuf"), "*.proto")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Type: schema.OriginProtobuf})
	}
	return ret, nil
}

func (p *ProtobufParser) IsValid(firstChars string) error {
	if !strings.Contains(firstChars, "syntax ") {
		return errors.New("not Protobuf")
	}
	return nil
}

func (p *ProtobufParser) log(rsp *ProtobufResponse, err error) {
	if err != nil {
		rsp.Rsp.Schema.Errors = append(rsp.Rsp.Schema.Errors, err.Error())
		p.logger.Error(fmt.Sprintf("unable to parse protobuf: %+v", err))
	}
}
