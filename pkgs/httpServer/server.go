package httpServer

import (
	"fmt"
	"net/http"
	"time"
)

func HttpServer() {

	called := 0

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		called++
		fmt.Fprint(w, called)
	})

	http.ListenAndServe(":8090", nil)
}
