package main

import (
	"go_dummy/main/apis"
	"go_dummy/main/core"
)

func main() {
	server := core.NewServer()
	apis.NewHealthyApi("/healthy", server)
	server.Start()
}
