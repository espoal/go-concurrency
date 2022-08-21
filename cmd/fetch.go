package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetch(responseChannel chan string) {
	resp, err := http.Get("http://localhost:8090/")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	responseChannel <- string(body)
}

func main() {

	concurrency := 10
	responseChannel := make(chan string, concurrency)
	n := 0
	for n < concurrency {
		go fetch(responseChannel)
		n++
	}

	n = 0
	for n < concurrency {
		fmt.Println(<-responseChannel)
		n++
	}

}
