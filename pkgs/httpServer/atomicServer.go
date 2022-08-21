package httpServer

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

func AtomicServer(parallelism int) {

	runtime.GOMAXPROCS(parallelism)

	var called uint32 = 0

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		value := atomic.AddUint32(&called, 1)
		//value := atomic.LoadUint32(&called)
		fmt.Fprint(w, value)
	})

	http.ListenAndServe(":8090", nil)
}
