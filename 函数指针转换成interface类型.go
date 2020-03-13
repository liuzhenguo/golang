//1. 如何将函数指针赋值给interface变量，进行转换
//2. 如何返回的类型和

package main

import (
	"fmt"
)

type Handler interface {
	// FrameWriter will be recycled when ServeQRPC finishes, so don't cache it
	ServeQRPC()
}
type handlerMW struct {
	t       int
	handler Handler
}
type HandlerFunc func()

func (f HandlerFunc) ServeQRPC() {

}
func (h *handlerMW) ServeQRPC() {
	fmt.Println("handlerMW")
}

//实现
func HandleFunc(handler func(), tt int) {
	Handle(HandlerFunc(handler), tt)
}

//调用
func Handle(handler Handler, tt int) {

	h := HandlerWithMW(handler, tt)
	h.ServeQRPC()
}

//返回
// HandlerWithMW wraps a handle with middleware
func HandlerWithMW(handler Handler, tt int) Handler {
	if tt == 0 {
		return handler
	}
	return &handlerMW{handler: handler, t: tt}
}

func main() {
	HandleFunc(func() {
		fmt.Println("hello world")
	}, 555)
}
