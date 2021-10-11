package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 20001,
	})
	if err != nil {
		fmt.Println("listen UDP failed,err:", err)
		return
	}
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("Read from UDP failed,err:", err)
			return
		}

		fmt.Println(data[:n])
		conn.WriteToUDP(data[:n], addr)
	}
}
