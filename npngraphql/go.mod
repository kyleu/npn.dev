module github.com/kyleu/npn/npngraphql
go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/graphql-go/graphql v0.7.9
	github.com/kyleu/npn/npnweb v0.0.0
	github.com/kyleu/npn/npncontroller v0.0.0
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npnconnection => ../npnconnection

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npncontroller => ../npncontroller

replace github.com/kyleu/npn/npnservice => ../npnservice

replace github.com/kyleu/npn/npntemplate => ../npntemplate

replace github.com/kyleu/npn/npnuser => ../npnuser

replace github.com/kyleu/npn/npnweb => ../npnweb
