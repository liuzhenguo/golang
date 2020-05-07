package control

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
	"zipkin/src/envbuild"
)

func invoall(wtu http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//多个请求超时控制并发请求
		var wg sync.WaitGroup
		fmt.Println("调用开始......")
		//设置超时时间
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(envbuild.GetTimeout()))

		//获取数据库的数据
		wg.Add(1)
		go func() {
			GetDbData(ctx)
			defer wg.Done()
		}()

		wg.Add(1)
		go func() {
			GetRedisData(ctx)
			defer wg.Done()
		}()

		wg.Add(1)
		go func() {
			GetRpcData(ctx)
			defer wg.Done()
		}()
		defer cancel()
		wg.Wait()
		wtu(w, r.WithContext(ctx))
	}
}

func wtouser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully"))
}

func Start() {
	http.Handle("/", invoall(wtouser))
	http.ListenAndServe("0.0.0.0:8899", nil)

}
