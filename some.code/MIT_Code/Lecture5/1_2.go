package main

import (
	"sync"
	"time"
)

func main() {

	counter := 0

	var mu sync.Mutex
	for i := 0; i < 10000; i++ {

		go func() {
			mu.Lock()
			counter = counter + 1
			mu.Unlock()
		}()

	}
	time.Sleep(1 * time.Second)
	mu.Lock()
	println(counter)
	mu.Unlock()
}
