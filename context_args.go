package main

import (
	"context"
	"fmt"
	"time"
)

func newconection(ctx context.Context) {
	ctx1, cancel := context.WithCancel(ctx)

	ctx1 = context.WithValue(ctx1, "k2", "v2")

	defer cancel()

	fmt.Println(ctx1.Value("k1").(string))
	fmt.Println(ctx1.Value("k2").(string))
}

func main() {

	fmt.Println("begin....")

	ctx, cancelctx := context.WithCancel(context.Background())

	ctx = context.WithValue(ctx, "k1", "v1")
	go newconection(ctx)

	time.Sleep(time.Second)

	fmt.Println("ctx:", ctx.Value("k1").(string))
	//fmt.Println("ctx:", ctx.Value("k2").(string))

	defer cancelctx()

	time.Sleep(time.Second)

}
