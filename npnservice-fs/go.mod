module github.com/kyleu/npn/npnservice-fs

go 1.15

require (
	github.com/kyleu/npn/npncore v1.0.0
	github.com/kyleu/npn/npnservice v1.0.0
	github.com/kyleu/npn/npnuser v1.0.0
)

replace github.com/kyleu/npn/npncore => ../npncore
replace github.com/kyleu/npn/npnservice => ../npnservice
replace github.com/kyleu/npn/npnuser => ../npnuser
