package main

import (
	"go-concurrency/v2/pkgs/httpClient"
	"go-concurrency/v2/pkgs/httpServer"
)

func main() {

	go httpClient.HttpClient(100, 11)

	httpServer.HttpServer(1)

}
