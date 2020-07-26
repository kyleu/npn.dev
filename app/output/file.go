package output

import (
	"encoding/json"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"strings"

	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/util"
)

type line struct {
	Indent  int
	Content string
}

type File struct {
	Pkg        util.Pkg
	Filename   string
	Type       FileType
	currIndent int
	output     []line
}

func NewGoFile(p *project.Project, pkg util.Pkg, key string) *File {
	if !strings.HasSuffix(key, ".go") {
		key += ".go"
	}
	if p != nil {
		pkg = append(p.RootPkg, pkg...)
	}
	return &File{Pkg: pkg, Filename: key, Type: FileTypeGo}
}

func (f *File) W(content string, indentDelta ...int) {
	for _, x := range indentDelta {
		if x < 0 {
			f.currIndent += x
		}
	}
	f.output = append(f.output, line{Indent: f.currIndent, Content: content})
	for _, x := range indentDelta {
		if x > 0 {
			f.currIndent += x
		}
	}
}

func (f *File) AddComment(content string) {
	l := f.comment(content)
	f.W(l.Content, l.Indent)
}

func (f *File) comment(content string) line {
	prefix := "# "
	switch f.Type {
	case FileTypeGo:
		prefix = "// "
	}
	return line{Content: prefix + content}
}

func (f *File) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})
	ret["path"] = strings.Join(append(f.Pkg, f.Filename), "/")
	ret["lines"] = f.Render()
	return json.Marshal(ret)
}

func (f *File) Render() []string {
	indent := "  "
	header := []line{f.comment(fmt.Sprintf("generated by [%v]", npncore.AppName))}
	var footer []line
	switch f.Type {
	case FileTypeGo:
		if len(f.Pkg) == 0 {
			header = append(header, line{Content: "package " + npncore.AppName})
		} else {
			header = append(header, line{Content: "package " + f.Pkg[len(f.Pkg)-1]})
		}
		header = append(header, line{})
		// indent = "\t"
	}
	lines := append(header, append(f.output, footer...)...)
	ret := make([]string, 0, len(lines))
	for _, l := range lines {
		indention := ""
		for i := 0; i < l.Indent; i++ {
			indention += indent
		}
		ret = append(ret, indention+l.Content)
	}
	return ret
}

func (f *File) Path() string {
	return f.Pkg.ToPath(f.Filename)
}
