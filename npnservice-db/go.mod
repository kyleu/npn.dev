module github.com/kyleu/npn/npnservice-db

go 1.15

require (
	github.com/kyleu/npn/npncore v1.0.0
	github.com/kyleu/npn/npndatabase v1.0.0
	github.com/kyleu/npn/npnservice v1.0.0
	github.com/kyleu/npn/npnuser v1.0.0
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
)

replace github.com/kyleu/npn/npncore => ../npncore
replace github.com/kyleu/npn/npndatabase => ../npndatabase
replace github.com/kyleu/npn/npnservice => ../npnservice
replace github.com/kyleu/npn/npnuser => ../npnuser
