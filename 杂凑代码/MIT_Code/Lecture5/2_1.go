//用于锁存在竞争关系，谁抢到谁执行，所以出现了
//alice和bob数据不正确.
//两个协程for中的操作不是原子性的存在抢占
//lock并不是保证锁定的数据不变
//去掉for循环的一对lock和unlock，for循环变成了原子性操作
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	alice := 10000
	bob := 10000

	var mu sync.Mutex

	total := alice + bob

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			mu.Unlock()
			mu.Lock()
			bob += 1
			mu.Unlock()
		}

	}()
	go func() {
		for i := 0; i < 1000; i++ {

			mu.Lock()
			bob -= 1
			mu.Unlock()

			mu.Lock()
			alice += 1
			mu.Unlock()
		}
	}()

	start := time.Now()

	for time.Since(start) < 3*time.Second {

		mu.Lock()
		if alice+bob != total {
			fmt.Printf("observed violation alice=%v bob=%v sum=%v\n", alice, bob, alice+bob)
		}
		mu.Unlock()

	}

}