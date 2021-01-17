module github.com/kyleu/npn/projects/wasm

go 1.15

require (
	emperror.dev/errors v0.8.0
	github.com/kyleu/npn v0.0.1
	github.com/kyleu/libnpn/npnconnection v0.0.1 // npn
	github.com/kyleu/libnpn/npncore v0.0.1 // npn
	github.com/kyleu/libnpn/npnweb v0.0.1 // npn
)

replace github.com/kyleu/npn => ../../

replace github.com/kyleu/libnpn/npnconnection => ../../libnpn/npnconnection

replace github.com/kyleu/libnpn/npncore => ../../libnpn/npncore

replace github.com/kyleu/libnpn/npnweb => ../../libnpn/npnweb
