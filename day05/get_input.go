package main

import (
	"bufio"
	"fmt"
	"os"
)

func userScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("您输入的内容是：%s\n", s)
}

func userBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("您输入的内容是：%s\n", s)
}

func main() {
	userScan()
}
