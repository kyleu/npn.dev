package parseintellij

import (
	"encoding/xml"
	"sort"
	"strings"

	"emperror.dev/errors"
	parseutil "github.com/kyleu/npn/app/model/parser/util"
	"github.com/kyleu/npn/app/util"
)

func (p *IntelliJParser) Parse(paths []string) (*parseutil.ParseResponse, error) {
	rsp, err := p.parse(paths, NewIntelliJResponse(paths))
	if err != nil {
		return nil, err
	}
	err = rsp.resolveChildren(1)
	if err != nil {
		return nil, errors.Wrap(err, "unable to resolve IntelliJ XML")
	}
	return rsp.Rsp, nil
}

func (p *IntelliJParser) parse(paths []string, ret *IntelliJResponse) (*IntelliJResponse, error) {
	rsp := ret
	var err error
	for _, pth := range paths {
		rsp, err = p.parsePath(pth, rsp)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing intelliJ")
		}
	}
	return rsp, nil
}

func (p *IntelliJParser) parsePath(path string, res *IntelliJResponse) (*IntelliJResponse, error) {
	unhandled := make(map[string]bool)
	onStart := func(e xml.StartElement, d *xml.Decoder) error {
		var err error
		switch e.Name.Local {
		case "dataSource":
			res.Rsp.Schema.Title = p.attrValue(e, util.KeyName)
			res.Rsp.Schema.Key = util.Slugify(res.Rsp.Schema.Title)
		case "database-model":
			// res.DBType = p.attrValue(e, "dbms")
			// res.DBFamily = p.attrValue(e, "family-id")
		case "root":
			err = res.extract(&ijRoot{}, e, d)
		case "database":
			err = res.extract(&ijDatabase{}, e, d)
		case "schema":
			err = res.extract(&ijSchema{}, e, d)
		case "sequence":
			err = res.extract(&ijSequence{}, e, d)
		case "object-type":
			err = res.extract(&ijObjectType{}, e, d)
		case "table":
			err = res.extract(&ijTable{}, e, d)
		case "column":
			err = res.extract(&ijColumn{}, e, d)
		case "index":
			err = res.extract(&ijIndex{}, e, d)
		case "key":
			err = res.extract(&ijKey{}, e, d)
		case "foreign-key":
			err = res.extract(&ijForeignKey{}, e, d)
		case "role", "access-method", "extension", "language", "operator", "routine", "argument":
			err = d.Skip()
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

func (p *IntelliJParser) attrValue(se xml.StartElement, name string) string {
	return parseutil.AttrValue(se, name, p.logger)
}
