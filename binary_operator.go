package main

import (
	"fmt"
)

func main() {

	length := 2147483648*2 - 1
	length1 := 256
	buf := make([]byte, 10)

	_ = append(buf[:0],
		byte(length>>24),
		byte(length>>16),
		byte(length>>8),
		byte(length),
	)
	fmt.Println(byte(length1))
	fmt.Println(buf)

}
