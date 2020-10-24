module github.com/kyleu/npn/npncontroller

go 1.15

require (
	github.com/kyleu/npn/npnconnection v0.0.17 // npn
	github.com/kyleu/npn/npncore v0.0.17 // npn
	github.com/kyleu/npn/npntemplate v0.0.17 // npn
	github.com/kyleu/npn/npnweb v0.0.17 // npn
)

replace github.com/kyleu/npn/npnconnection => ../npnconnection

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnservice => ../npnservice

replace github.com/kyleu/npn/npnuser => ../npnuser

replace github.com/kyleu/npn/npntemplate => ../npntemplate

replace github.com/kyleu/npn/npnweb => ../npnweb
