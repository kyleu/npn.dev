module github.com/kyleu/npn/npnconnection

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/kyleu/npn/npncore v1.0.0
	github.com/kyleu/npn/npnuser v1.0.0
	github.com/gorilla/websocket v1.4.1
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnuser => ../npnuser
