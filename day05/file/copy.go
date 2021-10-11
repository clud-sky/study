package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(dstName, srcName string) (write int64, err error) {
	// 以读的方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", srcName, err)
		return
	}

	// 关闭打开文件
	defer src.Close()

	// 以写｜创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dstName, err)
		return
	}

	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	_, err := copyFile("./dst.txt", "./xx.text")
	if err != nil {
		fmt.Println("copy file failed,err:", err)
		return
	}
	fmt.Println("copy done")
}
