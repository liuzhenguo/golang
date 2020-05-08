package main

import (
	"encoding/binary"
	"fmt"
)

func testBigEndian() {

	var testInt int32 = 256

	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(testInt))
	fmt.Println(bytes)
	covint := binary.BigEndian.Uint32(bytes)
	fmt.Println(covint)
}
func testLittleEndian() {
	var testint int32 = 256
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(testint))
	fmt.Println(bytes)

	convint := binary.LittleEndian.Uint32(bytes)
	fmt.Println(convint)
}
func main() {
	testBigEndian()
	fmt.Println("**********")
	testLittleEndian()
}
