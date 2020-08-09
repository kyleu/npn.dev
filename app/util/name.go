package util

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type Name struct {
	Imports []Pkg
	Pkg     Pkg
	Name    string
}

func (n Name) withPkg(s string) string {
	if len(n.Pkg) == 0 {
		return s
	}
	return n.Pkg[len(n.Pkg)-1] + "." + s
}

func (n Name) String() string {
	return n.withPkg(n.Name)
}

func (n Name) Class() string {
	return n.withPkg(strcase.ToCamel(n.Name))
}

func (n Name) Prop() string {
	return n.withPkg(strcase.ToLowerCamel(n.Name))
}

func (n Name) Plural() string {
	return inflection.Plural(n.Name)
}

func (n Name) PluralChoice(i int) string {
	if i == 1 {
		return n.Name
	}
	return n.Plural()
}

type NameRegistry struct {
	names        map[string]*Name
	replacements map[string]string
}

func NewNameRegistry(names map[string]*Name, replacements map[string]string) *NameRegistry {
	return &NameRegistry{names: names, replacements: replacements}
}

func (nr *NameRegistry) Get(path Pkg, name string, src Pkg) *Name {
	key := append(path, name).String()
	ret, ok := nr.names[key]
	if !ok {
		return &Name{Pkg: path.Trim(src), Name: name}
	}

	return ret
}

func (nr *NameRegistry) Set(path Pkg, name *Name) {
	nr.names[path.String()] = name
}

func (nr *NameRegistry) Replace(s string) string {
	for k, v := range nr.replacements {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
