package lib

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/npn/app/cli"
)

func Run() {
	port, err := cli.Run("0.0.0.0", 10101, "0.0.0", "master")
	if err != nil {
		panic(errors.WithStack(err))
	}
	println(fmt.Sprintf("npn library started on port [%v]", port))
}
