package main

import "fmt"

type Shap interface {
	Area() float32
}
type Circle struct {
	radius float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (cr *Circle) Area() float32 {

	return cr.radius * cr.radius
}
func main() {
	sq := new(Square)
	sq.side = 5

	var iarea Shap

	iarea = sq

	if t, ok := iarea.(*Square); ok {
		fmt.Printf("%T", t)
	}
	if c, ok := iarea.(*Circle); ok {
		fmt.Println( c)
	}
}
