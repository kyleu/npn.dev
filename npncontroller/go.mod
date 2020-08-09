module github.com/kyleu/npn/npncontroller

go 1.14

require (
	github.com/kyleu/npn/npncore v0.0.3
	github.com/kyleu/npn/npntemplate v0.0.2
	github.com/kyleu/npn/npnweb v0.0.2
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npndatabase => ../npndatabase

replace github.com/kyleu/npn/npnservice => ../npnservice

replace github.com/kyleu/npn/npnuser => ../npnuser

replace github.com/kyleu/npn/npntemplate => ../npntemplate

replace github.com/kyleu/npn/npnweb => ../npnweb
