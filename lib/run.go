package lib

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/cli"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
)

// Starts the application as a library, returning the actual TCP port the server is listening on
func Run(platform string, path string) int32 {
	cli.InitKeys()
	a := "0.0.0.0"
	p := uint16(0)
	info, r, err := cli.Start(platform, path)
	if err != nil {
		panic(errors.WithStack(err))
	}

	port, listener, err := npnweb.Listen(a, p)
	if err != nil {
		panic(errors.WithStack(err))
	}
	go func() {
		err = npnweb.Serve(listener, r)
		if err != nil {
			panic(errors.WithStack(err))
		}
	}()

	info.Logger().Info(fmt.Sprintf("%v library started using directory [%v] on port [%v]", npncore.AppName, info.Files().Root(), port))

	return int32(port)
}
