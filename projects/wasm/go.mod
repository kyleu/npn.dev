module github.com/kyleu/npn/projects/wasm

go 1.15

require (
	emperror.dev/errors v0.8.0
	github.com/kyleu/npn v0.0.21 // npn
	github.com/kyleu/npn/npnconnection v0.0.21 // npn
	github.com/kyleu/npn/npncontroller v0.0.21 // npn
	github.com/kyleu/npn/npncore v0.0.21 // npn
	github.com/kyleu/npn/npnuser v0.0.21 // npn
	github.com/kyleu/npn/npnweb v0.0.21 // npn
)

replace github.com/kyleu/npn => ../../

replace github.com/kyleu/npn/npnasset => ../../npnasset

replace github.com/kyleu/npn/npncontroller => ../../npncontroller

replace github.com/kyleu/npn/npnconnection => ../../npnconnection

replace github.com/kyleu/npn/npncore => ../../npncore

replace github.com/kyleu/npn/npndatabase => ../../npndatabase

replace github.com/kyleu/npn/npngraphql => ../../npngraphql

replace github.com/kyleu/npn/npnservice => ../../npnservice

replace github.com/kyleu/npn/npnservice-fs => ../../npnservice-fs

replace github.com/kyleu/npn/npnservice-db => ../../npnservice-db

replace github.com/kyleu/npn/npnscript => ../../npnscript

replace github.com/kyleu/npn/npnuser => ../../npnuser

replace github.com/kyleu/npn/npntemplate => ../../npntemplate

replace github.com/kyleu/npn/npnweb => ../../npnweb
