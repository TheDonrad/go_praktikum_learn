package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func thread(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&counter, 1)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {

		go thread(&wg)
	}
	wg.Wait()
	// программа должна выводить 200000
	fmt.Println(counter)
	ch := make(chan string, 1)
	ch <- "OK"
	fmt.Println(<-ch)
}
