package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//发送请求，获取连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err：", err)
		return
	}
	for {
		//输入数据并发送给服务器
		str := make([]byte, 4096)
		//读取键盘输入
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("os.Stdin.Read err：", err)
		}
		conn.Write(str[:n])
		if n == 0 {
			continue
		}
		n, err = conn.Read(str)
		if err == io.EOF {
			fmt.Println("已断开与服务器的连接")
			return
		}
		if err != nil {
			fmt.Println("os.Stdin.Read err：", err)
		}
		fmt.Print(string(str[:n]))
	}
}
