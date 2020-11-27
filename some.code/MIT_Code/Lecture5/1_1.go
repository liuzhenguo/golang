//多个go协程竞争同一个资源counter
package main

import (
	"time"
)

func main() {

	counter := 0

	for i := 0; i < 10000; i++ {

		go func() {
			counter = counter + 1
		}()

	}
	time.Sleep(1 * time.Second)
	println(counter)

}
