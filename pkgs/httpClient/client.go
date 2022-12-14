package httpClient

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
)

func fetch(responseChannel chan string) {
	resp, err := http.Get("http://localhost:8090/")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	text := string(body)
	if err != nil {
		panic(err)
	}

	responseChannel <- text
}

func HttpClient(concurrency int, parallelism int) {

	runtime.GOMAXPROCS(parallelism)

	countMap := make(map[string]bool, concurrency)
	responseChannel := make(chan string, concurrency)
	n := 0
	for n < concurrency {
		go fetch(responseChannel)
		n++
	}

	dataRaces := 0

	n = 0
	for n < concurrency {
		value := <-responseChannel
		if countMap[value] {
			dataRaces++
		} else {
			countMap[value] = true
		}
		n++
	}

	n = 1
	for n < concurrency+1 {
		if !countMap[strconv.Itoa(n)] {
			fmt.Println("Missing: ", n)
		}
		n++
	}

	fmt.Println("Data races: ", dataRaces)
	fmt.Println("Missing: ", dataRaces+len(countMap)-concurrency)

}
