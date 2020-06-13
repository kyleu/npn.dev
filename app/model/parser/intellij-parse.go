package parser

import (
	"emperror.dev/errors"
	"encoding/xml"
	"github.com/kyleu/npn/app/util"
	"sort"
	"strings"
)

func (r *IntelliJResponse) extract(x IntelliJResult, e xml.StartElement, d *xml.Decoder) error {
	err := d.DecodeElement(x, &e)
	if err != nil {
		return errors.Wrap(err, "Error decoding ["+e.Name.Local+"] item")
	}
	r.Data = append(r.Data, x)
	bp, _ := r.ByParent[x.ParentID()]
	bp = append(bp, x)
	r.ByParent[x.ParentID()] = bp
	return nil
}

func (p *IntelliJParser) ParseDataSourceXML(path string) (*IntelliJResponse, error) {
	res := NewIntelliJResponse()
	unhandled := make(map[string]bool)
	onStart := func(e xml.StartElement, d *xml.Decoder) error {
		var err error
		switch e.Name.Local {
		case "dataSource":
			res.DSN = p.attrValue(e, util.KeyName)
		case "database-model":
			res.DBType = p.attrValue(e, "dbms")
			res.DBFamily = p.attrValue(e, "family-id")
		case "root":
			err = res.extract(&ijRoot{}, e, d)
		case "database":
			err = res.extract(&ijDatabase{}, e, d)
		case "role":
			err = res.extract(&ijRole{}, e, d)
		case "schema":
			err = res.extract(&ijSchema{}, e, d)
		case "extension":
			err = res.extract(&ijExtension{}, e, d)
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
		case "access-method", "language", "operator", "routine", "argument":
			err = d.Skip()
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

func (p *IntelliJParser) attrValue(se xml.StartElement, name string) string {
	return attrValue(se, name, p.logger)
}
