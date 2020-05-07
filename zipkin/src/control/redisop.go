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
		fmt.Println("get redis value:", v)
	}()
	select {
	case <-ctx.Done():
		fmt.Println("其他调用失败")
		return ctx.Err()
	case e := <-err:
		fmt.Println("自己调用失败")
		return e
	default:
		return nil
	}
}
