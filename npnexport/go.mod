module github.com/kyleu/npn/npnexport

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/johnfercher/maroto v0.27.0
	github.com/kyleu/npn/npncore v0.0.14 // npn
)

replace github.com/kyleu/npn/npncore => ../npncore
