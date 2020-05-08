package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	ip := conn.RemoteAddr()
	fmt.Println("客户端已连接---", ip)
	//读取客户端数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("客户端断开连接")
			runtime.Goexit()
		}
		if err != nil {
			fmt.Println("conn.Read err：", err)
			return
		}
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {
			fmt.Println("客户端申请断开连接---", ip)
			fmt.Println("客户端已经断开连接---", ip)
			runtime.Goexit()
		}
		fmt.Print("已接收客户端输入：", string(buf[:n]))
		//将接收到的数据转换成大写，并发送给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
func main() {
	//创建服务
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	defer listener.Close()
	if err != nil {
		fmt.Println("net.Listen err：", err)
		return
	}
	//监听客户端请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err：", err)
			return
		}
		go handleConn(conn)
	}
}
