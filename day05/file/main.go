package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

	// 记得关闭文件
	defer fileObj.Close()

	// 读文件
	// var temp = make([]byte,128)
	var temp [128]byte
	for {
		n, err := fileObj.Read(temp[:])
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}
		//fmt.Println(n)
		fmt.Printf("读了%d个字节", n)
		fmt.Println(string(temp[:n]))
		if n < 0 {
			return
		}
	}
}

// 利用bufio这个包读取文件
func readFormFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		lile, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Print(lile)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed,err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	readFromFileByIoutil()
}
