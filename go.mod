module github.com/kyleu/npn

go 1.13

require (
	emperror.dev/emperror v0.32.0
	emperror.dev/errors v0.7.0
	emperror.dev/handler/logur v0.4.0
	github.com/cosmtrek/air v1.12.1 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/emicklei/proto v1.9.0
	github.com/fatih/color v1.9.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/sessions v1.2.0
	github.com/graph-gophers/graphql-go v0.0.0-20200309224638-dae41bde9ef9
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jinzhu/inflection v1.0.0
	github.com/kyleu/npn/npncore v0.0.1
	github.com/kyleu/npn/npndatabase v0.0.1
	github.com/kyleu/npn/npnuser v0.0.1
	github.com/kyleu/npn/npnweb v0.0.1
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pyros2097/go-embed v0.0.0-20200430035352-0689b0033fd1 // indirect
	github.com/sagikazarmark/ocmux v0.2.0
	github.com/santhosh-tekuri/jsonschema v1.2.4
	github.com/santhosh-tekuri/jsonschema/v2 v2.2.0
	github.com/shiyanhui/hero v0.0.2
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	golang.org/x/mobile v0.0.0-20200212152714-2b26a4705d24 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3
	golang.org/x/tools v0.0.0-20200708183856-df98bc6d456c // indirect
	logur.dev/adapter/logrus v0.4.1
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npncore => ./npncore
replace github.com/kyleu/npn/npndatabase => ./npndatabase
replace github.com/kyleu/npn/npnuser => ./npnuser
replace github.com/kyleu/npn/npnweb => ./npnweb
