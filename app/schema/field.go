package schema

import (
	"github.com/iancoleman/strcase"
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/schema/schematypes"
	"github.com/kyleu/npn/app/util"
)

type Field struct {
	Key      string              `json:"key"`
	Type     schematypes.Wrapped `json:"type"`
	Metadata *Metadata           `json:"metadata,omitempty"`
}

func (f *Field) String() string {
	return f.Key + " " + f.Type.String()
}

func (f *Field) StringFor(ft output.FileType, nr *util.NameRegistry, src util.Pkg) string {
	extra := ""
	if ft == output.FileTypeGo {
		extra = " `json:\"" + f.PropName(nr) + "\"`"
	}
	return f.ClassName(nr) + " " + f.Type.StringFor(ft, nr, src) + extra
}

func (f *Field) ClassName(nr *util.NameRegistry) string {
	return nr.Replace(strcase.ToCamel(f.Key))
}

func (f *Field) PropName(nr *util.NameRegistry) string {
	return nr.Replace(strcase.ToLowerCamel(f.Key))
}

type Fields []*Field

func (s Fields) Get(key string) *Field {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
