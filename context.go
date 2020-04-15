package main

import (
	"context"
	"fmt"
	"net/http"
)

/*
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

}*/

/*
1。函数指针的使用
2。interface的使用
3。函数回调。
*/
/*
type AA interface {
	ServerQ(a int, b int)
}
type BB func(int, int)

func (bb BB) ServerQ(a int, b int) {
	bb(a, b)
}
func Handler(test AA) {
	test.ServerQ(10, 20)
}
func Handle(test func(int, int)) {
	Handler(BB(test))
}
func main() {
	Handle(func(a int, b int) {
		fmt.Println(a, "a", b, "b")
	})

}*/
/*
type AA interface {
	ServerQ(a int, b int)
}

type BB struct {
	aa AA
	dd int
}

type CC func(a int, b int)

func (cc CC) ServerQ(a int, b int) {
	cc.ServerQ(a, b)
}
func (bb BB) ServerQ(a int, b int) {

}
func GetHandler() AA {
	var cc CC
	bb := BB{aa: cc, dd: 10}
	return bb
}

func main() {
	fmt.Println("begin..")
	GetHandler()

}
*/

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "key", "value")
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func handler(w http.ResponseWriter, req *http.Request) {

	value := req.Context().Value("key").(string)
	fmt.Println(value)
	//fmt.Fprintln(w, "value: ", value)
	fmt.Println("hello world")
	return
}

func main() {
	http.Handle("/", middleWare(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}
