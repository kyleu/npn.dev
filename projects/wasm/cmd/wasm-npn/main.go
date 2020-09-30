package main

import (
	"fmt"
	"syscall/js"

	"github.com/kyleu/npn/app/cli"
	"github.com/kyleu/npn/npncore"
)

var exitChannel = make(chan bool)

type Foo struct {
	X string
	Y int
	Z bool
}

func main() {
	println(fmt.Sprintf("Starting [%v] WASM", npncore.AppName))
	ai, _, err := cli.Start("wasm", ".")
	if err != nil {
		panic(err)
	}

	ai.Logger().Error("It's Alive!")

	// ret := &Foo{X: "asdf", Y: 110, Z: true}
	ret := "Hello, WASM!"

	js.Global().Set("npn", ret)

	<-exitChannel
}
