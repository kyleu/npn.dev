package npnweb

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncore"
)

func MakeServer(info AppInfo, r *mux.Router, address string, port uint16) (uint16, error) {
	var msg = "%v is starting on [%v:%v]"
	if info.Debug() {
		msg += " (verbose)"
	}
	info.Logger().Info(fmt.Sprintf(msg, npncore.AppName, address, port))
	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		return port, errors.Wrap(err, fmt.Sprintf("unable to listen on port [%v]", port))
	}
	if port == 0 {
		_, portStr := npncore.SplitString(l.Addr().String(), ':', true)
		actualPort, err := strconv.Atoi(portStr)
		if err != nil {
			return 0, errors.Wrap(err, "invalid port [" + portStr + "]")
		}
		if uint16(actualPort) != port {
			info.Logger().Info(fmt.Sprintf("started on http://localhost:%v", actualPort))
		}
		port = uint16(actualPort)
	}
	err = http.Serve(l, r)
	if err != nil {
		return port, errors.Wrap(err, "unable to run http server")
	}
	return port, nil
}
