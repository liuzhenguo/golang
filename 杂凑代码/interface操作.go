//代码很简单有几个需要注意的地方
//1.interface声明的变量接受的是一个指针，类似的还有map和切片本身就是指针类型
//2.实现了interface的方法可以赋值给interface的变量

package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	//var sq2 Square
	//sq2.side = 10
	//sq1 := new(Square)
	//sq1.side = 5

	var areaIntf Shaper
	//areaIntf = sq2
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
