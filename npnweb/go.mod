module github.com/kyleu/npn/npnweb

go 1.13

require (
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/sessions v1.2.0
	github.com/kyleu/npn/npncore v0.0.1
	github.com/sagikazarmark/ocmux v0.2.0
)

replace github.com/kyleu/npn/npncore => ../npncore
