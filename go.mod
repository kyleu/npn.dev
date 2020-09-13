module github.com/kyleu/npn

go 1.15

require (
	emperror.dev/emperror v0.32.0
	emperror.dev/errors v0.7.0
	emperror.dev/handler/logur v0.4.0
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.1
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/kyleu/npn/npnasset v1.0.0
	github.com/kyleu/npn/npnconnection v1.0.0
	github.com/kyleu/npn/npncontroller v1.0.0
	github.com/kyleu/npn/npncore v1.0.0
	github.com/kyleu/npn/npnservice v1.0.0
	github.com/kyleu/npn/npntemplate v1.0.0
	github.com/kyleu/npn/npnuser v1.0.0
	github.com/kyleu/npn/npnweb v1.0.0
	github.com/pyros2097/go-embed v0.0.0-20200430035352-0689b0033fd1 // indirect
	github.com/sagikazarmark/ocmux v0.2.0
	github.com/shiyanhui/hero v0.0.2
	github.com/spf13/cobra v0.0.5
	golang.org/x/mobile v0.0.0-20200801112145-973feb4309de // indirect
	golang.org/x/tools v0.0.0-20200913032122-97363e29fc9b // indirect
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npnasset => ./npnasset

replace github.com/kyleu/npn/npncontroller => ./npncontroller

replace github.com/kyleu/npn/npnconnection => ./npnconnection

replace github.com/kyleu/npn/npncore => ./npncore

replace github.com/kyleu/npn/npndatabase => ./npndatabase

replace github.com/kyleu/npn/npngraphql => ./npngraphql

replace github.com/kyleu/npn/npnservice => ./npnservice

replace github.com/kyleu/npn/npnuser => ./npnuser

replace github.com/kyleu/npn/npntemplate => ./npntemplate

replace github.com/kyleu/npn/npnweb => ./npnweb
