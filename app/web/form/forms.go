package form

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/mitchellh/mapstructure"
	"logur.dev/logur"
)

type ProfileForm struct {
	Theme     string `mapstructure:"theme"`
	LinkColor string `mapstructure:"linkColor"`
	NavColor  string `mapstructure:"navColor"`
	Ref       string `mapstructure:"ref"`
}

type ConnectionForm struct {
	Svc   string `mapstructure:"svc"`
	Cmd   string `mapstructure:"cmd"`
	Param string `mapstructure:"param"`
}

type SchemaSaveForm struct {
	Path     string `mapstructure:"path"`
	Key      string `mapstructure:"key"`
	Title    string `mapstructure:"title"`
}

func Decode(r *http.Request, tgt interface{}, logger logur.Logger) error {
	_ = r.ParseForm()
	frm := make(map[string]interface{}, len(r.Form))
	for k, v := range r.Form {
		frm[k] = strings.Join(v, "||")
	}
	md := &mapstructure.Metadata{}
	err := mapstructure.DecodeMetadata(frm, tgt, md)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to parse [%T] form", tgt))
	}

	if logger != nil {
		if len(md.Unused) > 0 {
			msg := fmt.Sprintf("parsed [%T] form with unused keys [%v]", tgt, strings.Join(md.Unused, ", "))
			logger.Warn(msg)
			bytes, _ := json.Marshal(tgt)
			logger.Warn(string(bytes))
		}
	}
	return nil
}
