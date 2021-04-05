package main

import (
	"arcstack/arcstack-chat-server/managers"
	"arcstack/arcstack-chat-server/pkg/setting"
)

func init() {
	setting.Setup()
}

func main() {
	chatManager := managers.InitialiseManager()
	chatManager.RunWsServer()
	return
}
