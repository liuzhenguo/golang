package main

import (
	"fmt"
)

type AA interface {
	Reader()
}
type BB struct {
	aa AA
	bb int
}

func (*BB) Reader() {

}
func Internal(aa AA) {

	aa.Reader()

}
func main() {

	b := new(BB)

	b.bb = 111111
	Internal(b)
	fmt.Println("hello world...")

}
