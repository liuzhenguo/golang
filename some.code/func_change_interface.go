/* 
	强转成handlerfunc后，赋值给接口，变成了handlerfunc类型，
	可以调用handlerfunc的方法，又因为本身又是一个函数可以直接
	调用自己

*/

package main

import (
	"fmt"
)

type Handler interface {
	HandleMessage()
}
type HandlerFunc func(a int) error

func (h HandlerFunc) HandleMessage() {
	fmt.Println("handlemessage!")
	h(10)
}

type Handle struct {
	a int
}

func (h Handle) handleMsg(a int) error {
	fmt.Println("a =", a)
	return nil
}

func AddHandler(h Handler) {
	h.HandleMessage()

}

func main() {
	h := Handle{}
	AddHandler(HandlerFunc(h.handleMsg))
}
