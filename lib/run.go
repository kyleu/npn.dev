package lib

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/app/cli"
)

func Run() {
	err := cli.Run("0.0.0.0", 10101, "0.0.0", "master")
	if err != nil {
		panic(errors.WithStack(err))
	}
}
