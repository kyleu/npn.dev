module github.com/kyleu/npn/npnweb

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/kyleu/npn/npncore v1.0.0
	github.com/kyleu/npn/npnservice v1.0.0
	github.com/kyleu/npn/npnuser v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/mitchellh/mapstructure v1.1.2
	github.com/gorilla/sessions v1.2.0
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npndatabase => ../npndatabase

replace github.com/kyleu/npn/npnservice => ../npnservice

replace github.com/kyleu/npn/npnuser => ../npnuser
