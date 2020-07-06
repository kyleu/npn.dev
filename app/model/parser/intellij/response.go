package parseintellij

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/kyleu/npn/app/model/schema/schematypes"
	"github.com/kyleu/npn/app/util"

	"emperror.dev/errors"
	parseutil "github.com/kyleu/npn/app/model/parser/util"
	"github.com/kyleu/npn/app/model/schema"
)

type IntelliJResult interface {
	ParentID() int
}

type IntelliJResponse struct {
	Rsp      *parseutil.ParseResponse
	ByParent map[int][]IntelliJResult
	state    util.Pkg
}

func NewIntelliJResponse(paths []string) *IntelliJResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginDatabase, Source: paths[0]}
	return &IntelliJResponse{
		Rsp:      parseutil.NewParseResponse(paths, md),
		ByParent: make(map[int][]IntelliJResult),
	}
}

func (r *IntelliJResponse) extract(x IntelliJResult, e xml.StartElement, d *xml.Decoder) error {
	err := d.DecodeElement(x, &e)
	if err != nil {
		return errors.Wrap(err, "Error decoding ["+e.Name.Local+"] item")
	}
	r.Rsp.Data = append(r.Rsp.Data, x)
	bp := r.ByParent[x.ParentID()]
	bp = append(bp, x)
	r.ByParent[x.ParentID()] = bp
	return nil
}

func (r *IntelliJResponse) resolveChild(child IntelliJResult) error {
	var err error
	switch c := child.(type) {
	case *ijDatabase:
		err = parseDatabase(r, c)
	case *ijSchema:
		err = parseSchema(r, c)
	case *ijObjectType:
		err = parseObjectType(r, c)
	case *ijTable:
		err = parseTable(r, c)
	case *ijColumn:
		err = parseColumn(r, c)
	case *ijIndex:
		err = parseIndex(r, c)
	case *ijKey:
		err = parseKey(r, c)
	case *ijForeignKey:
		err = parseForeignKey(r, c)
	case *ijSequence:
		parseSequence(r, c)
	default:
		err = errors.New(fmt.Sprintf("cannot resolve child of type [%T] (state: [%v])", child, strings.Join(r.state, ", ")))
	}
	return err
}

func (r *IntelliJResponse) resolveChildren(parentID int) error {
	children := r.ByParent[parentID]
	for _, child := range children {
		err := r.resolveChild(child)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseDatabase(r *IntelliJResponse, ij *ijDatabase) error {
	r.state = r.state.Push(ij.Name)
	err := r.resolveChildren(ij.ID)
	r.state = r.state.Shift()
	return err
}

func parseSchema(r *IntelliJResponse, ij *ijSchema) error {
	r.state = r.state.Push(ij.Name)
	err := r.resolveChildren(ij.ID)
	r.state = r.state.Shift()
	return err
}

func parseObjectType(r *IntelliJResponse, ij *ijObjectType) error {
	switch ij.SubKind {
	case "enum":
		labels := strings.Split(ij.Labels, "\n")
		for i, label := range labels {
			labels[i] = strings.TrimSpace(label)
		}
		fields := make(schema.Fields, 0, len(labels))
		for _, label := range labels {
			fields = append(fields, &schema.Field{
				Key:      label,
				Type:     schematypes.Wrap(schematypes.EnumValue{}),
				Metadata: nil,
			})
		}
		model := &schema.Model{Pkg: r.state, Key: ij.Name, Fields: fields, Type: schema.ModelTypeEnum}
		return r.Rsp.Schema.AddModel(model)
	default:
		return nil
	}
}

func parseTable(r *IntelliJResponse, ij *ijTable) error {
	m := &schema.Model{Pkg: r.state, Key: ij.Name, Type: schema.ModelTypeStruct}
	err := r.Rsp.Schema.AddModel(m)
	if err != nil {
		return err
	}
	r.state = r.state.Push(ij.Name)
	err = r.resolveChildren(ij.ID)
	r.state = r.state.Shift()
	return err
}

func parseColumn(r *IntelliJResponse, ij *ijColumn) error {
	table := r.Rsp.Schema.Models.Get(r.state.Shift(), r.state.Last())
	if table == nil {
		return errors.New("no table found with key [" + r.state.String() + "]")
	}
	t := parseutil.ParseDatabaseType(r.state.Shift(), ij.DataType, ij.NotNull == 0)
	err := table.AddField(&schema.Field{Key: ij.Name, Type: t})
	return err
}

func parseIndex(r *IntelliJResponse, ij *ijIndex) error {
	table := r.Rsp.Schema.Models.Get(r.state.Shift(), r.state.Last())
	if table == nil {
		return errors.New("no table found with key [" + r.state.String() + "]")
	}
	names := strings.Split(ij.ColNames, "\n")
	for idx, name := range names {
		names[idx] = strings.TrimSpace(name)
	}
	return table.AddIndex(&schema.Index{Key: ij.Name, Fields: names, Unique: ij.Unique != 0, Primary: ij.Primary != 0})
}

func parseKey(r *IntelliJResponse, ij *ijKey) error {
	r.state = r.state.Push(ij.Name)
	err := r.resolveChildren(ij.ID)
	r.state = r.state.Shift()
	return err
}

func parseForeignKey(r *IntelliJResponse, ij *ijForeignKey) error {
	r.state = r.state.Push(ij.Name)
	err := r.resolveChildren(ij.ID)
	r.state = r.state.Shift()
	return err
}

func parseSequence(r *IntelliJResponse, c *ijSequence) {
	r.Rsp.Data = append(r.Rsp.Data, c)
}
