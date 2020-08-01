module github.com/kyleu/npn/npntemplate

go 1.13

require (
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/kyleu/npn/npncore v0.0.1
	github.com/kyleu/npn/npnweb v0.0.1
	github.com/shiyanhui/hero v0.0.2
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npnweb => ../npnweb
