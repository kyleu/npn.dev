module github.com/kyleu/npn/npnconnection

go 1.15

require (
	emperror.dev/errors v0.7.0
	github.com/kyleu/npn/npncore v0.0.12 // npn
	github.com/kyleu/npn/npnuser v0.0.12 // npn
	github.com/gorilla/websocket v1.4.1
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnuser => ../npnuser
