module github.com/kyleu/npn/npnasset

go 1.15

require (
	emperror.dev/emperror v0.32.0
	emperror.dev/errors v0.7.0
	github.com/kyleu/npn/npnconnection v0.0.12 // npn
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnconnection => ../npnconnection

replace github.com/kyleu/npn/npnuser => ../npnuser
