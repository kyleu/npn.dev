package main

import (
	"fmt"
	"os"

	"github.com/kyleu/npn/app/cli"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
)

func main() {
	cmd := cli.Configure(version, commitHash)

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
