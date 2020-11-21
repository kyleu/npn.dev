package npnweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/mitchellh/mapstructure"
	"logur.dev/logur"
)

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

	if len(md.Unused) > 0 {
		msg := fmt.Sprintf("parsed [%T] form with unused keys [%v]", tgt, strings.Join(md.Unused, ", "))
		logger.Warn(msg)
		bytes, _ := json.Marshal(tgt)
		if logger == nil {
			fmt.Println(string(bytes))
		} else {
			logger.Warn(string(bytes))
		}
	}
	return nil
}
