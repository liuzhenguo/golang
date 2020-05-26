//channel中传递的数据对引用的修改
package main

import (
	"fmt"

	"time"
)

type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func (counter *Counter) String() string {
	return fmt.Sprintf("{count:%d}", counter.count)
}
func main() {

	sycChan := make(chan struct{}, 2)

	go func() {

		for {

			if elem, ok := <-mapChan; ok {

				counter := elem["count"]
				counter.count++

			} else {
				break
			}

		}
		fmt.Println("Stopped.[reciver]")
		sycChan <- struct{}{}

	}()

	go func() {

		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {

			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map:%v,[send]\n", countMap)
		}
		close(mapChan)
		sycChan <- struct{}{}

	}()
	<-sycChan
	<-sycChan
}
