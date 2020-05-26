package main

import (
	"fmt"
)

type myint int

func main() {

	mp1 := make(map[int][]*int)
	mp2 := make(map[int]*[]int)

	var a int = 10
	var b int = 20

	var b1 int = 10
	var b2 int = 20
	var t2 = []int{b1, b2}

	var tt = []*int{&a, &b}

	mp1[1] = tt
	mp2[1] = &t2
	fmt.Print(tt)
	fmt.Print(mp1)
	fmt.Print(mp2)

}
