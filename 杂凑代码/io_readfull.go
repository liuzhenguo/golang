package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	buf := make([]byte, 10, 10)
	buf2 := make([]byte, 10, 10)
	file, _ := os.Open("test.txt")
	n, err := io.ReadFull(file, buf)
	n2, _ := io.ReadFull(file, buf2)
	if err != nil {
		fmt.Println(n, n2, err.Error())
	} else {
		fmt.Println(string(buf))
		fmt.Println(string(buf2))
	}
}
