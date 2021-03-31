//main无法捕捉调用函数的panic,不同的协程无法捕捉panic
package main 

import "fmt"
import "sync"
import "time"
func main(){
		
	sg := sync.WaitGroup{}
	sg.Add(2)

	go func(){
		defer sg.Done()
		//panic("go1 panic")
		defer func(){
			if err:=recover();err!=nil{
				fmt.Println("panic:",err)
			}
		}()
		test()
	}()
		
	go func (){	
		defer sg.Done()
		time.Sleep(time.Second*3)	
	}()
	sg.Wait()
}
              
func test(){
	t1()
}
func t1(){
	t2()
}
func t2(){
	t3()
}
func t3(){
	panic("t3 panic!")
}
