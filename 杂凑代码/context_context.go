package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func newconection(ctx context.Context, wg *sync.WaitGroup) {
	ctx1, cancel := context.WithCancel(ctx)

	ctx1 = context.WithValue(ctx1, "k2", "v2")

	defer wg.Done()
	defer cancel()

	wg.Add(1)
	go newconection1(ctx1, wg)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx1 error:")
			return
		}

	}

}

func newconection1(ctx1 context.Context, wg *sync.WaitGroup) {

	ctx2, cancel := context.WithCancel(ctx1)
	ctx2 = context.WithValue(ctx2, "k3", "v3")
	defer cancel()
	defer wg.Done()

	select {
	case <-ctx2.Done():
		fmt.Println("ctx2 error:")
		return
	default:
		return
	}

}

func main() {

	fmt.Println("begin....")

	ctx, cancelctx := context.WithCancel(context.Background())

	ctx = context.WithValue(ctx, "k1", "v1")
	wg.Add(1)
	go newconection(ctx, &wg)

	//fmt.Println("ctx:", ctx.Value("k1").(string))

	cancelctx()

	wg.Wait()

}
