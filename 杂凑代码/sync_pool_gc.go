package main

import (
	"fmt"
	//"runtime"
	"sync"
)

type TestPool struct {
	a int
	s string
}

func TestPoolFunc(a int, s string) {
	tpool := &sync.Pool{New: func() interface{} { return &TestPool{a: 10, s: "hi"} }}

	ts := tpool.Get().(*TestPool)
	ts.a = a
	ts.s = s

	tpool.Put(ts)

	//runtime.GC()
	ts = tpool.Get().(*TestPool)

	fmt.Println(ts.a, ts.s)
}
func main() {
	TestPoolFunc(100, "hello world")
}
