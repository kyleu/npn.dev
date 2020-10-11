module github.com/kyleu/npn/npntemplate

go 1.15

require (
	cloud.google.com/go v0.65.0 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/kyleu/npn/npnconnection v0.0.16 // npn
	github.com/kyleu/npn/npncore v0.0.16 // npn
	github.com/kyleu/npn/npnservice v0.0.16 // npn
	github.com/kyleu/npn/npnuser v0.0.16 // npn
	github.com/kyleu/npn/npnweb v0.0.16 // npn
	github.com/shiyanhui/hero v0.0.2
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npnconnection => ../npnconnection

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnuser => ../npnuser

replace github.com/kyleu/npn/npnservice => ../npnservice

replace github.com/kyleu/npn/npnweb => ../npnweb
