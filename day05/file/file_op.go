package main

import (
	"fmt"
	"io"
	"os"
)

// 文件中间插入
func file1() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	// 因为没有办法直接再文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed,err:%v", err)
		return
	}

	// 读取文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed,err:%v", err)
		return
	}

	// 写入临时文件
	tmpFile.Write(ret[:n])
	// 再写入要插入的内容
	var s []byte
	s = []byte("e")
	tmpFile.Write(s)
	// 紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}

	// 源文件的也写入临时文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")
}

func main() {
	file1()
	var phone = "13689050228"
	fmt.Printf(phone[:9])
}
