package parseintellij

import (
	"emperror.dev/errors"
	"encoding/xml"
	"fmt"
	parseutil "github.com/kyleu/npn/app/model/parser/util"
	"github.com/kyleu/npn/app/model/schema"
	"strings"
)

type IntelliJResult interface {
	ParentID() int
}

type IntelliJResponse struct {
	RootFile string `json:"root"`
	DBType   string
	DBFamily string
	Data     []IntelliJResult
	ByParent map[int][]IntelliJResult
	Schema   *schema.Schema `json:"schema"`
	state    map[string]string
}

func NewIntelliJResponse(paths []string) *IntelliJResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginIntelliJ, Source: paths[0]}
	return &IntelliJResponse{
		RootFile: paths[0],
		Data:     make([]IntelliJResult, 0),
		ByParent: make(map[int][]IntelliJResult),
		Schema:   schema.NewSchema(paths[0], paths, &md),
		state:    make(map[string]string),
	}
}

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

func (r *IntelliJResponse) resolve() error {
	children, ok := r.ByParent[1]
	if !ok {
		return errors.New("no children for root")
	}
	for _, child := range children {
		err := r.resolveChild(child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *IntelliJResponse) resolveChild(child IntelliJResult) error {
	var err error
	switch c := child.(type) {
	case *ijRole:
		// noop: err = parseRole(r, c)
	case *ijDatabase:
		err = parseDatabase(r, c)
	case *ijExtension:
		err = parseExtension(r, c)
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
		err = parseSequence(r, c)
	default:
		stateStrings := make([]string, 0, len(r.state))
		for k, v := range r.state {
			stateStrings = append(stateStrings, k+": "+v)
		}
		err = errors.New(fmt.Sprintf("cannot resolve child of type [%T] (state: [%v])", child, strings.Join(stateStrings, ", ")))
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
	r.state["database"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "database")
	return err
}

func parseExtension(r *IntelliJResponse, ij *ijExtension) error {
	r.state["extension"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "extension")
	return err
}

func parseSchema(r *IntelliJResponse, ij *ijSchema) error {
	r.state["schema"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "schema")
	return err
}

func parseObjectType(r *IntelliJResponse, ij *ijObjectType) error {
	r.state["objectType"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "objectType")
	return err
}

func parseTable(r *IntelliJResponse, ij *ijTable) error {
	r.state["table"] = ij.Name
	err := r.Schema.AddModel(&schema.Model{Key: ij.Name, Type: schema.ModelTypeStruct})
	if err != nil {
		return err
	}
	err = r.resolveChildren(ij.ID)
	delete(r.state, "table")
	return err
}

func parseColumn(r *IntelliJResponse, ij *ijColumn) error {
	table := r.Schema.Models.Get(nil, r.state["table"])
	if table == nil {
		return errors.New("no table found with key [" + r.state["table"] + "]")
	}
	t := parseutil.ParseDatabaseType(ij.DataType, ij.NotNull == 0)
	err := table.AddField(&schema.Field{Key: ij.Name, Type: t})
	return err
}

func parseIndex(r *IntelliJResponse, ij *ijIndex) error {
	r.state["index"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "index")
	return err
}

func parseKey(r *IntelliJResponse, ij *ijKey) error {
	r.state["key"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "key")
	return err
}

func parseForeignKey(r *IntelliJResponse, ij *ijForeignKey) error {
	r.state["foreignKey"] = ij.Name
	err := r.resolveChildren(ij.ID)
	delete(r.state, "foreignKey")
	return err
}

func parseSequence(r *IntelliJResponse, c *ijSequence) error {
	r.Data = append(r.Data, c)
	return nil
}
