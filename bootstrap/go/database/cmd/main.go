package main

import (
	"fmt"
	"os"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
)

func main() {
	err := start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() error {
	fmt.Println(fmt.Sprintf("{{.Key}} [%v:%v]", version, commitHash))
	return nil
}
