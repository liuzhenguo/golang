package main

import (
	"fmt"
)

type TestA interface {
	TestFunc()
}

type TotalS struct {
	ln TestA
}

type TestB struct {
	TestA

	B string
}

func (bb *TestB) TestFunc() {
	fmt.Println(bb.B)
}

type TestC struct {
	TestA
	C string
}

func (cc *TestC) TestFunc() {
	fmt.Println(cc.C)
}

type TestD struct {
	TestA
	D string
}

func (dd *TestD) TestFunc() {
	fmt.Println(dd.D)
}

func main() {

	var bb TestB
	var cc TestC
	var dd TestD

	bb.B = "bbbbb"
	cc.C = "ccccc"
	dd.D = "ddddd"

	var ts1, ts2, ts3 TotalS
	var ts []TotalS
	ts1.ln = &bb
	ts2.ln = &cc
	ts3.ln = &dd

	ts = append(ts, ts1)
	ts = append(ts, ts2)
	ts = append(ts, ts3)

	for _, val := range ts {
		val.ln.TestFunc()
	}

}
