package main

import (
	"fmt"
	"os"
)

// 获取命令行参数
func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%T\n", os.Args)
	fmt.Println(os.Args[1], os.Args[2])
}
