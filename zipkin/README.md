##### zipkin demo 

1. 需要配置数据库和redis，redis和db的配置在etc/config.ini文件中
2. 配置完毕后下载需要安装的库，使用的是zipkin-go
3. 运行（最好能打包到docker中，比较懒不想折腾了）


```
cd test
docker-compose up -d
sudo su
echo -e "\n127.0.0.1 zipkin" >> /etc/hosts
cd ../src && go build -o zipkin.bin
./zipkin.bin
curl 127.0.0.1:8899/handler
```



注：目前只实现了监测一个http请求,目前可参考的资料比较少，demo还不够完善。