package main

import "fmt"
import "time"


func main(){
	
	timer := time.NewTimer(time.Second*2)
	ch := make(chan int,10)
	go func(){

		in :=0
		
		for{
			in++
			ch<-in
			
		}

	}()

	select{
		case <-ch:
		case <-timer.C
			fmt.Println("这样调用")
		//case <-time.After(time.Second):c//会产生内存泄漏！！！！

	}

}
