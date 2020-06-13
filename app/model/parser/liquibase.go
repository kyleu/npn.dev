package parser

import (
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"logur.dev/logur"
	"path"
)

type LiquibaseParser struct {
	Key     string
	logger  logur.Logger
}

func (p *LiquibaseParser) Detect(root string) ([]schema.DataSource, error) {
	fs, err := util.GetMatchingFiles(path.Join(root, "data", "liquibase"), "*.xml")
	if err != nil {
		return nil, err
	}
	ret := make([]schema.DataSource, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, schema.DataSource{Key: f, Paths: []string{f}, Origin: schema.OriginLiquibase})
	}
	return ret, nil
}
