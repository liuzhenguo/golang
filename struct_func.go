package main

import (
	"fmt"
)

type AA struct {
}

func (aa *AA) TestA() {

	fmt.Println("test one")
}

type BB struct {
	AA
	a int
}

func main() {

	bb := BB{}

	bb.TestA()
}
