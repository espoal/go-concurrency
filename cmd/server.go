package main

import "go-concurrency/v2/pkgs/httpServer"

func main() {

	httpServer.HttpServer(1)
	//httpServer.AtomicServer(6)
}
