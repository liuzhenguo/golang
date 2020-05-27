package main

import (
    "sync"
    "fmt"
    "runtime"
    "time"
)

func main() {
    p := &sync.Pool{
        New: func() interface{} {
            return 0
        },
    }

    runtime.GOMAXPROCS(2)

    a := p.Get().(int)
    fmt.Println(a)
    p.Put(1)

    wg := sync.WaitGroup{}
    wg.Add(1)
    go func(){
        defer wg.Done()
        p.Put(100)
    }()
    wg.Wait()

    time.Sleep(time.Second * 1)

    p.Put(4)
    p.Put(5)

    fmt.Println(p.Get())
    fmt.Println(p.Get())
    fmt.Println(p.Get())
        // fmt.Println(p.Get())
}
