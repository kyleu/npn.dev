package main

import (
	"fmt"
	"os"

	"github.com/kyleu/npn/app/cli"
)

func main() {
	cmd := cli.Configure()

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
