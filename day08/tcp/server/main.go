package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端
func processConn(conn net.Conn) {
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))

		fmt.Print("请回复：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	// 1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start server 127.0.0.1:20000 on failed,err:", err)
		return
	}
	// 2、等待别人来更我建立链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		// 3、与客户端通信
		go processConn(conn)
	}
}
