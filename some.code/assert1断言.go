package main
import "fmt"


func main(){
	
	switch t :=areaInf.(type){
		case *Square:
			fmt.Println("square")
		case *Circle:
			fmt.Println("circle")
		default:
			fmt.Println("unkown!")
	}

}
