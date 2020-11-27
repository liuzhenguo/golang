package main

import (
	"fmt"
	"time"

	"github.com/oklog/run"
)
func main() {

	g := run.Group{}

	g.Add(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("hello 22222222")
		return nil
	}, func(err error) {
		fmt.Println(err)
	})
	g.Add(func() error {
		fmt.Println("hello 111111")
		return nil
	}, func(err error) {
		fmt.Println("hello error deal with.. 1111")
	})

	err := g.Run()
	fmt.Println(err)
}