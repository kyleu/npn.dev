package main

import (
	"fmt"
	"syscall/js"

	"github.com/kyleu/libnpn/npnconnection"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/cli"
)

var exitChannel = make(chan bool)
var callback *js.Value
var ai npnweb.AppInfo

func register(_ js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return "must provide 1 argument: callback"
	}
	callback = &args[0]

	cli.FileLoaderOverride = NewLocalStorageLoader()

	var err error
	ai, _, err = cli.Start("wasm", ".")
	if err != nil {
		panic(err)
	}

	log(fmt.Sprintf("Started [%v] in bypass mode", npncore.AppName))

	svc := app.Svc(ai)
	svc.Socket.SetWASMCallback(func(s string) { callback.Invoke(s) })

	return nil
}

func handle(_ js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		logErr("must provide one argument containing message data")
		return nil
	}
	msgString := args[0].String()

	if callback == nil {
		logErr("nil callback")
	}

	msg := &npnconnection.Message{}
	err := npncore.FromJSONStrict([]byte(msgString), msg)
	if err != nil {
		logErr("invalid JSON [" + msgString + "]: " + err.Error())
		return nil
	}

	svc := app.Svc(ai)
	err = npnconnection.OnMessage(svc.Socket, npnconnection.WASMID, msg)
	if err != nil {
		logErr("error processing message [" + msgString + "]: " + err.Error())
		return nil
	}

	// log(fmt.Sprintf("handling [%v]", msgString))
	return nil
}

func shutdown(_ js.Value, _ []js.Value) interface{} {
	exitChannel <- true
	return nil
}

func main() {
	js.Global().Set("npn_register", js.FuncOf(register))
	js.Global().Set("npn_handler", js.FuncOf(handle))
	js.Global().Set("npn_shutdown", js.FuncOf(shutdown))

	log("X")
	<-exitChannel
	log("Y")

	log("run complete")
}

func logErr(s string) {
	log("ERROR: " + s)
}

func log(s string) {
	println("[WASM] " + s)
}
