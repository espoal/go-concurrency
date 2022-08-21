package main

import (
	"go-concurrency/v2/pkgs/httpServer"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)

	httpServer.HttpServer()
}
