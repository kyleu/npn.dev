module github.com/kyleu/npn/npnasset

require (
	emperror.dev/emperror v0.32.0
	emperror.dev/errors v0.7.0
	github.com/kyleu/npn/npnconnection v1.0.0
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnconnection => ../npnconnection

replace github.com/kyleu/npn/npnuser => ../npnuser

go 1.15
