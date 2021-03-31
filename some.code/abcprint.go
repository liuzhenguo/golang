package  main

import (
	"fmt"
	"sync"
)


var cha chan string
var chb chan string
var chc chan string
var sg sync.WaitGroup
func init() {

	cha = make(chan string)
	chb = make(chan string)
	chc = make(chan string)
	sg = sync.WaitGroup{}
}

func main(){
	sg.Add(3)
	go processa()
	go processb()
	go processc()
	sg.Wait()


}


func processa(){
	for i :=0;i<10;i++{
		cha <-"A"
		fmt.Println(<-chc)

	}
	defer sg.Done()


}
func processb(){
	for i:=0;i<10;i++{
		fmt.Println(<-cha)
		chb<-"B"

	}

	defer sg.Done()

}
func processc(){
	for i:=0;i<10;i++{
		fmt.Println(<-chb)
		chc<-"C"
	}

	defer sg.Done()
}

