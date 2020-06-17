package parseprotobuf

import (
	"fmt"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"path"
)

type ProtobufParser struct {
	Key     string
	logger  logur.Logger
}

func NewParser(logger logur.Logger) *ProtobufParser {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: schema.OriginProtobuf.Key})
	return &ProtobufParser{Key: schema.OriginProtobuf.Key, logger: logger}
}

func (p *ProtobufParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := util.GetMatchingFiles(path.Join(root, "data", "protobuf"), "*.proto")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Origin: schema.OriginProtobuf})
	}
	return ret, nil
}

func (p *ProtobufParser) log(rsp *ProtobufResponse, err error) {
	if err != nil {
		rsp.Schema.Errors = append(rsp.Schema.Errors, err.Error())
		p.logger.Error(fmt.Sprintf("unable to parse protobuf: %+v", err))
	}
}
