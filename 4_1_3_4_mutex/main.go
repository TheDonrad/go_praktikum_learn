package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[int]int)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(v int) {
			mu.Lock()
			m[v] = v
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}
