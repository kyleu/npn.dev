package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting WASM test server on port 10100")
	err := http.ListenAndServe("0.0.0.0:10100", http.FileServer(http.Dir("assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
