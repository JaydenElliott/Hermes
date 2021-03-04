package main

import (
	"arcstack/arcstack-chat-server/managers"
)

func main() {
	chatManager := managers.InitialiseManager()
	chatManager.RunWsServer()
	return
}
