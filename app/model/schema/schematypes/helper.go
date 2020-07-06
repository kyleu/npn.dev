package schematypes

import "github.com/kyleu/npn/app/util"

func ErrorWrapped(err error) Wrapped {
	return Wrap(Error{Message: err.Error()})
}

func ErrorString(msg string) Wrapped {
	return Wrap(Error{Message: msg})
}

func OptionWrapped(t Type) Wrapped {
	return Wrap(Option{T: Wrap(t)})
}

func ReferenceWrapped(pkg util.Pkg, t string) Wrapped {
	return Wrap(Reference{Pkg: pkg, T: t})
}

var NilWrapped = Wrap(Nil{})
var StringWrapped = Wrap(String{})
