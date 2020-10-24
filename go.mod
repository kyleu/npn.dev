module github.com/kyleu/npn

go 1.15

replace github.com/kyleu/npn/npnasset => ./npnasset

replace github.com/kyleu/npn/npnconnection => ./npnconnection

replace github.com/kyleu/npn/npncontroller => ./npncontroller

replace github.com/kyleu/npn/npncore => ./npncore

replace github.com/kyleu/npn/npndatabase => ./npndatabase

replace github.com/kyleu/npn/npnexport => ./npnexport

replace github.com/kyleu/npn/npngraphql => ./npngraphql

replace github.com/kyleu/npn/npnqueue => ./npnqueue

replace github.com/kyleu/npn/npnscript => ./npnscript

replace github.com/kyleu/npn/npnservice => ./npnservice

replace github.com/kyleu/npn/npnservice-db => ./npnservice-db

replace github.com/kyleu/npn/npnservice-fs => ./npnservice-fs

replace github.com/kyleu/npn/npnuser => ./npnuser

replace github.com/kyleu/npn/npntemplate => ./npntemplate

replace github.com/kyleu/npn/npnweb => ./npnweb

require (
	emperror.dev/emperror v0.33.0
	emperror.dev/errors v0.8.0
	emperror.dev/handler/logur v0.4.0
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/kirsle/configdir v0.0.0-20170128060238-e45d2f54772f
	github.com/kyleu/npn/npnasset v0.0.17 // npn
	github.com/kyleu/npn/npnconnection v0.0.17 // npn
	github.com/kyleu/npn/npncontroller v0.0.17 // npn
	github.com/kyleu/npn/npncore v0.0.17 // npn
	github.com/kyleu/npn/npndatabase v0.0.0-00010101000000-000000000000 // indirect
	github.com/kyleu/npn/npnservice v0.0.17 // npn
	github.com/kyleu/npn/npnservice-fs v0.0.17 // npn
	github.com/kyleu/npn/npntemplate v0.0.17 // npn
	github.com/kyleu/npn/npnuser v0.0.17 // npn
	github.com/kyleu/npn/npnweb v0.0.17 // npn
	github.com/mccutchen/go-httpbin v1.1.1
	github.com/sagikazarmark/ocmux v0.2.0
	github.com/shiyanhui/hero v0.0.2
	github.com/spf13/cobra v1.0.0
	golang.org/x/text v0.3.3
	logur.dev/logur v0.17.0
)
