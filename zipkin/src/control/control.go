package control

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
	"zipkin/src/envbuild"

	"github.com/gorilla/mux"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

var tracer *zipkin.Tracer

func invoall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//return func(w http.ResponseWriter, r *http.Request) {
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
		// wg.Add(1)
		// go func() {
		// 	GetRpcData(ctx)
		// 	defer wg.Done()
		// }()
		defer cancel()
		wg.Wait()
		w.Write([]byte("successfully"))
	}

}
func invohandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var wg sync.WaitGroup

		fmt.Println("调用开始......")
		//设置超时时间
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(envbuild.GetTimeout()))

		//将span放在ctx中
		//ctx = zipkin.NewContext(ctx, span)

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
		go func() {
			nextfunc(w, r)
		}()
		go func() {
			nextone(w, r)
		}()
		defer cancel()
		wg.Wait()
		w.Write([]byte("successfully"))
	}

}
func nextfunc(w http.ResponseWriter, r *http.Request) {

	span, _ := tracer.StartSpanFromContext(r.Context(), "nextfunc")

	span.Tag("nextfunc111", "sleep 100 ms")

	defer span.Finish()
	time.Sleep(time.Millisecond * 100)
	span.Annotate(time.Now(), "expensive_call_func")
}
func nextone(w http.ResponseWriter, r *http.Request) {

	span, _ := tracer.StartSpanFromContext(r.Context(), "nextone")
	span.Tag("nextone", "sleep 150 ms")
	defer span.Finish()

	time.Sleep(time.Millisecond * 150)
	span.Annotate(time.Now(), "expensive_call_nexone_end")

}
func Start() {

	tracer = gettracer("zipkin_trace_101", "0.0.0.0:8899")
	serverMiddleware := zipkinhttp.NewServerMiddleware(
		tracer, zipkinhttp.TagResponseSize(true),
	)
	//路由
	router := mux.NewRouter()
	//监测接口
	router.Methods("GET").Path("/").HandlerFunc(invoall())
	router.Methods("GET").Path("/handler").HandlerFunc(invohandler())
	http.ListenAndServe("0.0.0.0:8899", serverMiddleware(router))

}
