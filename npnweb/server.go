package npnweb

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncore"
)

func MakeServer(info AppInfo, r *mux.Router, address string, port uint16) error {
	var msg = "%v is starting on [%v:%v]"
	if info.Debug() {
		msg += " (verbose)"
	}
	info.Logger().Info(fmt.Sprintf(msg, npncore.AppName, address, port))
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), r)
	return errors.Wrap(err, "unable to run http server")
}
