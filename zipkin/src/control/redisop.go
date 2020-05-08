package control

import (
	"context"
	"fmt"
	"zipkin/src/envbuild"

	"github.com/garyburd/redigo/redis"
)

func GetRedisData(ctx context.Context) error {
	//操作redis获取数据
	pool := envbuild.GetRedisPool()

	chstr := make(chan string)

	err := make(chan error)
	go func() {
		conn := pool.Get()
		defer conn.Close()
		key := fmt.Sprintf("String:Activeid:LimitUsers:183815")
		v, err := redis.String(conn.Do("GET", key))

		if err != nil {
			fmt.Println("redis invoke error", err)
			return
		}
		if len(v) == 0 {
			fmt.Println("redis nill return")
		}
		chstr <- v
	}()
	select {
	case <-ctx.Done():
		fmt.Println("redis其他调用失败:", ctx.Err())
		return ctx.Err()
	case e := <-err:
		fmt.Println("redis自己调用失败")
		return e
	case result := <-chstr:
		fmt.Println("redis data:", result)
		return nil
	}
}
