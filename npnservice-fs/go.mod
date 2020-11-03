module github.com/kyleu/npn/npnservice-fs

go 1.15

require (
	github.com/kyleu/npn/npncore v0.0.22 // npn
	github.com/kyleu/npn/npnservice v0.0.22 // npn
	github.com/kyleu/npn/npnuser v0.0.22 // npn
)

replace github.com/kyleu/npn/npncore => ../npncore
replace github.com/kyleu/npn/npnservice => ../npnservice
replace github.com/kyleu/npn/npnuser => ../npnuser
