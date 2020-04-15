package main

import (
	"fmt"
)

type SeverMutx struct {
}

func (mtx *SeverMutx) Hander() {

}
func (mtx *SeverMutx) Handle(cmd int, handler Handler) {
	handler.ServerQRPC(10, 20)
}

//程序执行的地方
type Handler interface {
	ServerQRPC(a int, b int)
}
type HandlerFunc func(a int, b int)

func (f HandlerFunc) ServerQRPC(a int, b int) {
	f(a, b)
}
func (mtx *SeverMutx) HandleFunc(cmd int, handler func(a int, b int)) {
	mtx.Handle(cmd, HandlerFunc(handler))
}

func NewFunction() *SeverMutx {

	return new(SeverMutx)

}
func main() {
	servermutx := NewFunction()

	servermutx.HandleFunc(1, func(a int, b int) {
		fmt.Println("a", a, "b", b)
	})

}
