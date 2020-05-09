module github.com/liuzhenguo/golang/zipkin/src

go 1.14

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.7.4
	github.com/msbranco/goconfig v0.0.0-20160629072055-3189001257ce
	github.com/openzipkin/zipkin-go v0.2.2
	zipkin/src v0.0.0-00010101000000-000000000000
)

replace zipkin/src => ../src
