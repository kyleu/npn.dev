package lib

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/npn/app/cli"
	"github.com/kyleu/npn/npnweb"
)

func Run() int32 {
	a := "0.0.0.0"
	p := uint16(10101)
	info, r, err := cli.Start("0.0.0", "master")
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

	info.Logger().Info(fmt.Sprintf("npn library started on port [%v]", port))

	return int32(port)
}
