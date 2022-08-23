package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	r := make([]int, 1<<20)
	runtime.GOMAXPROCS(1)
	for i := 0; i < 1<<20; i++ {
		wg.Add(1)
		go func(i int) {
			r[i] = i // race here
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(r))
}
