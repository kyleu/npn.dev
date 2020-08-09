package export

import (
	"fmt"
	"github.com/kyleu/npn/app/output"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/npncore"
	"strings"
)

func WriteGo(file *output.File, model *schema.Model, nr *util.NameRegistry) {
	switch model.Type {
	case schema.ModelTypeEnum:
		writeGoEnum(file, model, nr)
	default:
		writeGoStruct(file, model, nr)
	}
}

func writeGoEnum(file *output.File, model *schema.Model, nr *util.NameRegistry) {
	mcl := model.ClassName(nr)
	file.W("type "+mcl+" struct {", 1)
	file.W("Key string `json:\"key\"`")
	file.W("}", -1)
	file.W("")
	file.W("var (", 1)
	for _, field := range model.Fields {
		name := mcl + field.ClassName(nr)
		file.W(fmt.Sprintf("%v = %v{Key: \"%v\"}", name, mcl, field.Key))
	}
	file.W(fmt.Sprintf("%vUnknown = %v{Key: \"unknown\"}", mcl, mcl))
	file.W(")", -1)
	file.W("")
	var names []string
	for _, field := range model.Fields {
		names = append(names, mcl + field.ClassName(nr))
	}
	file.W("var All" + npncore.Plural(mcl) + " = []" + mcl + "{" + strings.Join(names, ", ") + "}")
	file.W("")

	file.W("func " + mcl + "FromString(s string) " + mcl + " {", 1)
	file.W("for _, t := range All" + npncore.Plural(mcl) + " {", 1)
	file.W("if t.Key == s {", 1)
	file.W("return t")
	file.W("}", -1)
	file.W("}", -1)
	file.W("return " + mcl + "Unknown")
	file.W("}", -1)
	file.W("")

	file.W("func (t *" + mcl + ") String() string {", 1)
	file.W("return t.Key")
	file.W("}", -1)
	file.W("")

	file.W("func (t *" + mcl + ") MarshalJSON() ([]byte, error) {", 1)
	file.W("return json.Marshal(t.Key)")
	file.W("}", -1)
	file.W("")

	file.W("func (t *" + mcl + ") UnmarshalJSON(data []byte) error {", 1)

	file.W("var s string")
	file.W("err := json.Unmarshal(data, &s)")
	file.W("if err != nil {", 1)
	file.W("return err")
	file.W("}", -1)

	file.W("*t = " + mcl + "FromString(s)")
	file.W("return nil")
	file.W("}", -1)
}

func writeGoStruct(file *output.File, model *schema.Model, nr *util.NameRegistry) {
	file.W("type "+model.ClassName(nr)+" struct {", 1)
	for _, field := range model.Fields {
		file.W(field.StringFor(output.FileTypeGo, nr, model.Pkg))
	}
	file.W("}", -1)
}
