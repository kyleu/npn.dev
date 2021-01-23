module github.com/kyleu/npn

go 1.15

replace github.com/kyleu/libnpn/npnconnection => ./libnpn/npnconnection

replace github.com/kyleu/libnpn/npncontroller => ./libnpn/npncontroller

replace github.com/kyleu/libnpn/npncore => ./libnpn/npncore

replace github.com/kyleu/libnpn/npnservice => ./libnpn/npnservice

replace github.com/kyleu/libnpn/npnservice-fs => ./libnpn/npnservice-fs

replace github.com/kyleu/libnpn/npnuser => ./libnpn/npnuser

replace github.com/kyleu/libnpn/npntemplate => ./libnpn/npntemplate

replace github.com/kyleu/libnpn/npnweb => ./libnpn/npnweb

require (
	emperror.dev/emperror v0.33.0
	emperror.dev/errors v0.8.0
	emperror.dev/handler/logur v0.4.0
	github.com/andybalholm/brotli v1.0.1
	github.com/getkin/kin-openapi v0.33.0
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/kirsle/configdir v0.0.0-20170128060238-e45d2f54772f
	github.com/kyleu/libnpn/npnconnection v0.0.1 // npn
	github.com/kyleu/libnpn/npncontroller v0.0.1 // npn
	github.com/kyleu/libnpn/npncore v0.0.48 // npn
	github.com/kyleu/libnpn/npnservice v0.0.1 // npn
	github.com/kyleu/libnpn/npnservice-fs v0.0.1 // npn
	github.com/kyleu/libnpn/npntemplate v0.0.1 // npn
	github.com/kyleu/libnpn/npnuser v0.0.48 // npn
	github.com/kyleu/libnpn/npnweb v0.0.1 // npn
	github.com/rbretecher/go-postman-collection v0.3.0
	github.com/sagikazarmark/ocmux v0.2.0
	github.com/shiyanhui/hero v0.0.2
	github.com/spf13/cobra v1.0.0
	golang.org/x/text v0.3.3
	logur.dev/logur v0.17.0
)
