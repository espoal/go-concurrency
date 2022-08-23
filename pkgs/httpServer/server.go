package httpServer

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func HttpServer(parallelism int) {

	runtime.GOMAXPROCS(parallelism)

	called := 0

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		called++
		temp := called
		fmt.Fprint(w, temp)
	})

	http.ListenAndServe(":8090", nil)
}
