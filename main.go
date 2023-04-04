package main

import (
	"go_dummy/main/apis"
	"go_dummy/main/core"
)

func main() {
	server := core.NewServer()
	_ = apis.NewHealthyApi("/", server)
	server.Start()
}
