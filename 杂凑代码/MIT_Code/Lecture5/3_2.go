//for  循环会将一个核心打满
package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex

	//cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {

		go func() {
			vote := requestVote()

			mu.Lock()
			defer mu.Unlock()

			if vote {
				count++
			}
			finished++
			//cond.Broadcast()

		}()

	}

	for {

		//cond.Wait()
		mu.Lock()

		if count >= 5 || finished == 10 {
			break
		}
		mu.Unlock()

	}
	if count >= 5 {
		println("recive 5+ votes!")
	} else {
		println("lost")
	}

}
func requestVote() bool {

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
