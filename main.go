package main

import (
	"arcstack/arcstack-chat-server/managers"
	"flag"
	"fmt"
	"net/http"
)

// Temp function for testing
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving home ", w, r)
}

// Temp address for testing put in config
var addr = flag.String("addr", ":8080", "http service address")

func main() {

	chatManager := managers.InitialiseManager()
	go chatManager.RunWsServer()
	http.HandleFunc("/", serveHome)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println("ListenAndServe Error: ", err)
	}
	return
}
