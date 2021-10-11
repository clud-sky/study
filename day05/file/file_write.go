package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDome() {
	fileObj, err := os.OpenFile("./xx.text", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()
	// write
	fileObj.Write([]byte("周林懵逼了\n"))
	fileObj.WriteString("周林解释不了")
	fileObj.Close()
}

func writeDome1() {
	fileObj, err := os.OpenFile("./xx.text", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()

	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello望见了") //写到缓存中
	wr.Flush()                 // 将缓存中的内容写入文件
}

func writeDome2() {
	str := "hello收到您能独立房间里"
	err := ioutil.WriteFile("./xx.text", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}

// 打开文件写入内容
func main() {
	writeDome2()
}
