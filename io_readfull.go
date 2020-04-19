package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:9900")

	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			break
		}

		go handle(conn)

	}

}
func handle(conn net.Conn) {
	alldatas := make([]byte, 10)
	rf := bufio.NewReader(conn)

	for {
		tmpbytes := make([]byte, 16)
		n, err := io.ReadFull(rf, tmpbytes)

		if err != io.EOF {
			alldatas = append(alldatas, tmpbytes[:n]...)
		} else {
			if err != nil && n == 0 {
				fmt.Println("no bytes!")
				break
			} else {
				fmt.Println("errror:", err)
				break
			}
		}

	}
	fmt.Println(string(alldatas))

	conn.Write([]byte("successfully"))
	defer conn.Close()
}
