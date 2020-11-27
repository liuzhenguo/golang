package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("server has been start===>")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":50001")
	//服务器端一般不定位具体的客户端套接字
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	ConnMap := make(map[string]*net.TCPConn)

	for {
		conn, _ := tcpListener.AcceptTCP()

		wg.Add(1)

		go func() {
			//读取数据
			reader := bufio.NewReader(conn)

			for {
				line, err := reader.ReadString('\n')

				if err != nil {
					break
				}

				fmt.Println(line)
			}
			fmt.Println("read all...")
			ConnMap[conn.RemoteAddr().String()] = conn
			fmt.Println("连接的客户端信息：", conn.RemoteAddr().String())
			defer func() {
				wg.Done()
				conn.Close()

			}()
		}()

	}
}
