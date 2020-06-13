package parser

import (
	"encoding/xml"
	"sort"
	"strings"
)

func (p *LiquibaseParser) ParseChangeLogXML(path string) (*LiquibaseResponse, error) {
	res := NewLiquibaseResponse()
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

	err := parseXML(path, onStart)
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
	return attrValue(se, name, p.logger)
}
