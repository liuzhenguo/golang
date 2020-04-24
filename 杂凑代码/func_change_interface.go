package main

import (
	"fmt"
)

type HandlerFunc func(a int) error

type Handle struct {
	a int
}

func (h *Handle) HandlerMessage(b int) error {
	fmt.Println("b = ", b)
	return nil
}
func AddHandler(h HandlerFunc) {
	h(10)
}
func main() {
	h := Handle{}
	AddHandler(HandlerFunc(h.HandlerMessage))
}
