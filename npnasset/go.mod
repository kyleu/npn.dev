module github.com/kyleu/npn/npnasset

go 1.14

require (
	emperror.dev/emperror v0.32.0
	emperror.dev/errors v0.7.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kyleu/npn/npncontroller v1.0.0
	github.com/kyleu/npn/npnweb v1.0.0
	github.com/sirupsen/logrus v1.4.2 // indirect
	golang.org/x/text v0.3.2 // indirect
	logur.dev/adapter/logrus v0.4.1 // indirect
	logur.dev/logur v0.16.2 // indirect
)


replace github.com/kyleu/npn/npnconnection => ../npnconnection

replace github.com/kyleu/npn/npncontroller => ../npncontroller

replace github.com/kyleu/npn/npncore => ../npncore

replace github.com/kyleu/npn/npndatabase => ../npndatabase

replace github.com/kyleu/npn/npnservice => ../npnservice

replace github.com/kyleu/npn/npnuser => ../npnuser

replace github.com/kyleu/npn/npntemplate => ../npntemplate

replace github.com/kyleu/npn/npnweb => ../npnweb
