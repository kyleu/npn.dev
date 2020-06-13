package parser

import (
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"path"
)

type IntelliJParser struct {
	Key     string
	logger  logur.Logger
}

func (p *IntelliJParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := util.GetMatchingFiles(path.Join(root, ".idea", "dataSources"), "*.xml")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Origin: schema.OriginIntelliJ})
	}
	return ret, nil
}
