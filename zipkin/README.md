#####1 需要配置数据库和redis，redis和db的配置在etc/config.ini文件中
#####2.配置完毕后下载需要安装的库，使用的是zipkin-go
#####3.运行（最好能打包到docker中，比较懒不想折腾了）

####注：目前只实现了监测一个http请求,在想监听的地方使用span，按照官网zipkin-go的例程进行埋点。