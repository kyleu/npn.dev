package main

import (
	"fmt"

	"github.com/webview/webview"
)

func main() {
	port := 10101
	launchWebview(port)
}

func launchWebview(port int) {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("npn")
	w.SetSize(1280, 720, webview.HintNone)
	w.Navigate(fmt.Sprintf("http://localhost:%v", port))
	w.Run()
}
