package parseliquibase

import (
	"emperror.dev/errors"
	"encoding/xml"
	parseutil "github.com/kyleu/npn/app/model/parser/util"
	"sort"
	"strings"
)

func (p *LiquibaseParser) ParseChangeLogXML(paths []string) (*LiquibaseResponse, error) {
	return p.parse(paths, NewLiquibaseResponse(paths))
}

func (p *LiquibaseParser) parse(paths []string, ret *LiquibaseResponse) (*LiquibaseResponse, error) {
	rsp := ret
	var err error
	for _, pth := range paths {
		rsp, err = p.parsePath(pth, rsp)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing liquibase")
		}
	}
	return rsp, nil
}

func (p *LiquibaseParser) parsePath(path string, res *LiquibaseResponse) (*LiquibaseResponse, error) {
	unhandled := make(map[string]bool)
	onStart := func(e xml.StartElement, d *xml.Decoder) error {
		var err error
		switch e.Name.Local {
		case "databaseChangeLog":
			// noop
		case "changeSet":
			// noop
		case "createTable":
			err = res.extract(&lCreateTable{}, e, d)
		case "createIndex":
			err = res.extract(&lCreateIndex{}, e, d)
		case "addUniqueConstraint":
			err = res.extract(&lAddUniqueConstraint{}, e, d)
		case "addForeignKeyConstraint":
			err = res.extract(&lAddForeignKeyConstraint{}, e, d)
		default:
			unhandled[e.Name.Local] = true
			err = d.Skip()
		}
		return err
	}

	err := parseutil.ParseXML(path, onStart)
	if err != nil {
		return nil, err
	}
	if len(unhandled) > 0 {
		u := make([]string, 0, len(unhandled))
		for k := range unhandled {
			u = append(u, k)
		}
		sort.Strings(u)
		uStr := strings.Join(u, ", ")
		p.logger.Warn("unhandled elements [" + uStr + "] from XML at [" + path + "]")
	}
	return res, nil
}

func (p *LiquibaseParser) attrValue(se xml.StartElement, name string) string {
	return parseutil.AttrValue(se, name, p.logger)
}
